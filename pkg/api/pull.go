package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
)

type PullsOption struct {
	EnableComments bool
	EnableReviews  bool
	EnableCommits  bool
	EnableStatuses bool
	Rules          *PullRequestRules
}

func (c *Client) GetPulls(ctx context.Context, owner string, repo string, opt PullsOption) ([]*PullRequest, error) {
	pagenation := NewPagenation()
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

	pulls := make([]*PullRequest, 0)
	pullsIndexesByNumber := make(map[int]int)
	pullsIndexesBySHA := make(map[string]int)
	for i := 0; i < pagenation.RequestedNum(); i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			goto L
		case res := <-pagenation.ch:
			if err := res.Error(); err != nil {
				return nil, err
			}
			switch body := res.Interface().(type) {
			case []*github.PullRequest:
				values := make([]*PullRequest, len(body))
				for i, v := range body {
					index := len(pulls) + i
					values[i] = newPullRequest(owner, repo, v, nil, nil, nil, nil)
					pullsIndexesByNumber[v.GetNumber()] = index
					pullsIndexesBySHA[v.GetHead().GetSHA()] = index
				}
				pulls = append(pulls, values...)

				for _, v := range body {
					go func(pull *github.PullRequest) {
						if opt.EnableComments {
							pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								listCommentOptions := &github.PullRequestListCommentsOptions{
									Sort:        "created",
									Direction:   "desc",
									ListOptions: *opt,
								}
								return c.github.PullRequests.ListComments(childCtx, owner, repo, pull.GetNumber(), listCommentOptions)
							})
						}
						if opt.EnableReviews {
							pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.PullRequests.ListReviews(childCtx, owner, repo, pull.GetNumber(), opt)
							})
						}
						if opt.EnableCommits {
							pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.PullRequests.ListCommits(childCtx, owner, repo, pull.GetNumber(), opt)
							})
						}
						if opt.EnableStatuses {
							pagenation.Request(childCtx, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.Repositories.ListStatuses(childCtx, owner, repo, pull.GetHead().GetSHA(), opt)
							})
						}
					}(v)
				}
			case []*github.PullRequestComment:
				// /repos/[owner]/[repo]/pulls/[pr number]/comments
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				index := pullsIndexesByNumber[number]
				comments := make([]*PullRequestComment, len(body))
				for i, comment := range body {
					comments[i] = newPullRequestComment(comment)
				}
				pulls[index].Comments = append(pulls[index].Comments, comments...)
			case []*github.PullRequestReview:
				// /repos/[owner]/[repo]/pulls/[pr number]/reviews
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				index := pullsIndexesByNumber[number]
				reviews := make([]*PullRequestReview, len(body))
				for i, review := range body {
					reviews[i] = newPullRequestReview(review)
				}
				pulls[index].Reviews = append(pulls[index].Reviews, reviews...)
			case []*github.RepositoryCommit:
				// /repos/[owner]/[repo]/pulls/[pr number]/commits
				s := strings.Split(res.Response().Request.URL.Path, "/")
				number, _ := strconv.Atoi(s[len(s)-2])
				index := pullsIndexesByNumber[number]
				commits := make([]*RepositoryCommit, len(body))
				for i, commit := range body {
					commits[i] = newRepositoryCommit(commit)
				}
				pulls[index].Commits = append(pulls[index].Commits, commits...)
			case []*github.RepoStatus:
				// /repos/[owner]/[repo]/commits/[sha]/statuses
				s := strings.Split(res.Response().Request.URL.Path, "/")
				sha := s[len(s)-2]
				index := pullsIndexesBySHA[sha]
				statuses := make([]*RepoStatus, len(body))
				for i, status := range body {
					statuses[i] = newRepoStatus(status)
				}
				pulls[index].Statuses = append(pulls[index].Statuses, statuses...)
			}
		}
	}
L:

	return opt.Rules.Apply(pulls)
}
