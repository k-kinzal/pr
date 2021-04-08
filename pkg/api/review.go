package api

import (
	"context"
	"github.com/google/go-github/v28/github"
	"golang.org/x/sync/errgroup"
)

type ReviewOption struct {
	Action string
}

func (c *Client) AddApproval(ctx context.Context, pulls []*PullRequest, opt *ReviewOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				pullRequestReviewRequest := &github.PullRequestReviewRequest{Event: github.String(opt.Action)}
				pullRequestReview, _, err := c.github.PullRequests.CreateReview(ctx, pull.Owner, pull.Repo, int(pull.Number), pullRequestReviewRequest)
				if err != nil {
					return err
				}
				pull.Reviews = append(pull.Reviews, newPullRequestReview(pullRequestReview))
				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}
