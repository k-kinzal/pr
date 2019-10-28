package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
)

type PullRequest struct {
	*github.PullRequest
	Owner    *string                      `json:"-"`
	Repo     *string                      `json:"-"`
	Comments []*github.PullRequestComment `json:"comments"`
	Reviews  []*github.PullRequestReview  `json:"reviews"`
	Commits  []*github.RepositoryCommit   `json:"commits"`
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
	pagenation := new(Pagenation)
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	if num := opt.Rules.GetNumber(); num > 0 {
		go func() {
			pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
				return c.github.PullRequests.Get(childCtx, owner, repo, num)
			})
		}()
	} else {
		pagenation.RequestWithLimit(childCtx, opt.Rules.GetLimit()/PerPage, func(listOptions *github.ListOptions) (interface{}, *github.Response, error) {
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
	commitsMap := make(map[int][]*github.RepositoryCommit)
	statusesMap := make(map[string][]*github.RepoStatus)
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			goto L
		case res, ok := <-pagenation.Done():
			if !ok {
				goto L
			}
			if err := res.Error(); err != nil {
				return nil, err
			}
			switch v := res.Interface().(type) {
			case *github.PullRequest:
				pulls = append(pulls, v)
				pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
					listCommentOptions := &github.PullRequestListCommentsOptions{
						Sort:        "created",
						Direction:   "desc",
						ListOptions: *opt,
					}
					return c.github.PullRequests.ListComments(childCtx, owner, repo, v.GetNumber(), listCommentOptions)
				})
				pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
					return c.github.PullRequests.ListReviews(childCtx, owner, repo, v.GetNumber(), opt)
				})
				pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
					return c.github.PullRequests.ListCommits(childCtx, owner, repo, v.GetNumber(), opt)
				})
				pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
					return c.github.Repositories.ListStatuses(childCtx, owner, repo, v.GetHead().GetSHA(), opt)
				})
			case *github.PullRequestComment:
				// /repos/[owner]/[repo]/pulls/[pr number]/comments
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				commentsMap[number] = append(commentsMap[number], v)
			case *github.PullRequestReview:
				// /repos/[owner]/[repo]/pulls/[pr number]/reviews
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				reviewsMap[number] = append(reviewsMap[number], v)
			case *github.RepositoryCommit:
				// /repos/[owner]/[repo]/pulls/[pr number]/commits
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				commitsMap[number] = append(commitsMap[number], v)
			case *github.RepoStatus:
				// /repos/[owner]/[repo]/commits/[sha]/statuses
				s := strings.Split(res.Response().Request.URL.Path, "/")
				sha := s[len(s)-2]
				statusesMap[sha] = append(statusesMap[sha], v)
			}
		}
	}
L:

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
		commits, ok := commitsMap[pull.GetNumber()]
		if !ok {
			commits = make([]*github.RepositoryCommit, 0)
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
			Commits:     commits,
			Statuses:    statuses,
		})
	}

	return opt.Rules.Apply(allPulls)
}
