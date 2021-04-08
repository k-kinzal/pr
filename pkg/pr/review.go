package pr

import (
	"context"
	"fmt"
	"github.com/k-kinzal/pr/pkg/api"
	"strings"
)

const (
	ReviewActionAddApproval = "APPROVE"
)

type ReviewOption struct {
	Action string
	*ListOption
}

func Review(owner string, repo string, opt *ReviewOption) ([]*api.PullRequest, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOption := &api.Options{
		Token:     token,
		RateLimit: opt.Rate,
	}
	client := api.NewClient(ctx, clientOption)

	pullOption := api.PullsOption{
		EnableComments: opt.EnableComments,
		EnableReviews:  opt.EnableReviews,
		EnableCommits:  opt.EnableCommits,
		EnableStatuses: opt.EnableStatuses,
		EnableChecks:   opt.EnableChecks,
		Rules:          api.NewPullRequestRules(opt.Rules, opt.Limit),
	}
	pulls, err := client.GetPulls(ctx, owner, repo, pullOption)
	if err != nil {
		return nil, err
	}
	if len(pulls) == 0 {
		return nil, &NoMatchError{pullOption.Rules}
	}

	// GitHub Pull Request Reviews require `event` parameter to be uppercase
	// https://docs.github.com/en/rest/reference/pulls#create-a-review-for-a-pull-request
	action := strings.ToUpper(opt.Action)

	switch action {
	case ReviewActionAddApproval:
		reviewOption := &api.ReviewOption{
			Action: action,
		}
		return client.AddApproval(ctx, pulls, reviewOption)
	}
	return nil, fmt.Errorf("currently, `approve` is only supported action, but %s was passed", opt.Action)
}
