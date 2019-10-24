package api

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
)

const (
	PerPage = 100
)

type PullRequest struct {
	*github.PullRequest
	Owner    *string                      `json:"-"`
	Repo     *string                      `json:"-"`
	Comments []*github.PullRequestComment `json:"comments"`
	Reviews  []*github.PullRequestReview  `json:"reviews"`
	Statuses []*github.RepoStatus         `json:"statuses"`
}

func (pr *PullRequest) GetOwner() string {
	if pr.Owner == nil {
		return ""
	} else {
		return *pr.Owner
	}
}

func (pr *PullRequest) GetRepo() string {
	if pr.Repo == nil {
		return ""
	} else {
		return *pr.Repo
	}
}

type PullsOption struct {
	Rules *PullRequestRules
}

func (c *Client) GetPulls(ctx context.Context, owner string, repo string, opt PullsOption) ([]*PullRequest, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	requestNum := 0
	getAllPage := func(ch chan interface{}, errCh chan error, limit int, callback func(listOptions *github.ListOptions) (interface{}, *github.Response, error)) {
		requestNum++
		go func() {
			listOptions := &github.ListOptions{
				Page:    0,
				PerPage: PerPage,
			}
			data, resp, err := callback(listOptions)
			if err != nil {
				errCh <- err
				return
			}
			ch <- data
			if resp.NextPage == -1 {
				return
			}
			allPage := resp.LastPage
			if limit > 0 && allPage > limit/listOptions.PerPage {
				allPage = limit / listOptions.PerPage
			}
			for i := 1; i < allPage; i++ {
				requestNum++
				go func(page int) {
					listOptions.Page = page
					data, _, err := callback(listOptions)
					if err != nil {
						errCh <- err
						return
					}
					ch <- data
				}(i)
			}
		}()
	}

	ch := make(chan interface{})
	errCh := make(chan error)

	if num := opt.Rules.GetNumber(); num > 0 {
		requestNum++
		go func() {
			pull, _, err := c.github.PullRequests.Get(childCtx, owner, repo, num)
			if err != nil {
				errCh <- err
				return
			}
			ch <- []*github.PullRequest{pull}
		}()
	} else {
		getAllPage(ch, errCh, opt.Rules.GetLimit(), func(listOptions *github.ListOptions) (interface{}, *github.Response, error) {
			pullOptions := &github.PullRequestListOptions{
				State:       opt.Rules.GetState(),
				Head:        opt.Rules.GetHead(),
				Base:        opt.Rules.GetBase(),
				Sort:        "created",
				Direction:   "desc",
				ListOptions: *listOptions,
			}
			return c.github.PullRequests.List(childCtx, owner, repo, pullOptions)
		})
	}

	pulls := make([]*github.PullRequest, 0)
	commentsMap := make(map[int][]*github.PullRequestComment)
	reviewsMap := make(map[int][]*github.PullRequestReview)
	statusesMap := make(map[string][]*github.RepoStatus)
	for i := 0; i < requestNum; i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			break
		case data := <-ch:
			switch reflect.TypeOf(data).Kind() {
			case reflect.Slice, reflect.Array:
				v := reflect.ValueOf(data)
				for i := 0; i < v.Len(); i++ {
					switch vv := v.Index(i).Interface().(type) {
					case *github.PullRequest:
						pulls = append(pulls, vv)
						func(pull *github.PullRequest) {
							getAllPage(ch, errCh, -1, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								listCommentOptions := &github.PullRequestListCommentsOptions{
									Sort:        "created",
									Direction:   "desc",
									ListOptions: *opt,
								}
								return c.github.PullRequests.ListComments(childCtx, owner, repo, pull.GetNumber(), listCommentOptions)
							})
							getAllPage(ch, errCh, -1, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.PullRequests.ListReviews(childCtx, owner, repo, pull.GetNumber(), opt)
							})
							getAllPage(ch, errCh, -1, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.Repositories.ListStatuses(childCtx, owner, repo, pull.GetHead().GetSHA(), opt)
							})
						}(vv)
					case *github.PullRequestComment:
						// https://api.github.com/repos/[owner]/[repo]/pulls/[pr number]
						s := strings.Split(vv.GetPullRequestURL(), "/")
						number, _ := strconv.Atoi(s[len(s)-1])
						commentsMap[number] = append(commentsMap[number], vv)
					case *github.PullRequestReview:
						// https://api.github.com/repos/[owner]/[repo]/pulls/[pr number]
						s := strings.Split(vv.GetPullRequestURL(), "/")
						number, _ := strconv.Atoi(s[len(s)-1])
						reviewsMap[number] = append(reviewsMap[number], vv)
					case *github.RepoStatus:
						// https://api.github.com/repos/[owner]/[repo]/statuses/[sha]
						s := strings.Split(vv.GetURL(), "/")
						statusesMap[s[len(s)-1]] = append(statusesMap[s[len(s)-1]], vv)
					}
				}
			}
		case err := <-errCh:
			cancel()
			return nil, err
		}
	}

	var allPulls []*PullRequest
	for _, pull := range pulls {
		comments, ok := commentsMap[pull.GetNumber()]
		if !ok {
			comments = make([]*github.PullRequestComment, 0)
		}
		reviews, ok := reviewsMap[pull.GetNumber()]
		if !ok {
			reviews = make([]*github.PullRequestReview, 0)
		}
		statuses, ok := statusesMap[pull.GetHead().GetSHA()]
		if !ok {
			statuses = make([]*github.RepoStatus, 0)
		}
		allPulls = append(allPulls, &PullRequest{
			PullRequest: pull,
			Owner:       &owner,
			Repo:        &repo,
			Comments:    comments,
			Reviews:     reviews,
			Statuses:    statuses,
		})
	}

	return opt.Rules.Apply(allPulls)
}
