package api

import (
	"context"
	"reflect"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/google/go-github/v28/github"
)

const (
	PerPage = 100
)

type PullRequest struct {
	*github.PullRequest
	Owner    *string              `json:"-"`
	Repo     *string              `json:"-"`
	Statuses []*github.RepoStatus `json:"statuses"`
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
	Rules     *PullRequestRules
	RateLimit int
}

func (c *Client) GetPulls(ctx context.Context, owner string, repo string, opt PullsOption) ([]*PullRequest, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	rateLimit := opt.RateLimit
	requestNum := 0
	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(rateLimit)), rateLimit)
	getAllPage := func(ch chan interface{}, errCh chan error, limit int, callback func(listOptions *github.ListOptions) (interface{}, *github.Response, error)) {
		requestNum++
		go func() {
			if err := limiter.Wait(childCtx); err != nil {
				errCh <- err
				return
			}
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
					if err := limiter.Wait(childCtx); err != nil {
						errCh <- err
						return
					}
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
								return c.github.Repositories.ListStatuses(childCtx, owner, repo, pull.GetHead().GetSHA(), opt)
							})
						}(vv)
					case *github.RepoStatus:
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
		statuses, ok := statusesMap[pull.GetHead().GetRef()]
		if !ok {
			statuses = make([]*github.RepoStatus, 0)
		}
		allPulls = append(allPulls, &PullRequest{
			PullRequest: pull,
			Owner:       &owner,
			Repo:        &repo,
			Statuses:    statuses,
		})
	}

	return opt.Rules.Apply(allPulls)
}
