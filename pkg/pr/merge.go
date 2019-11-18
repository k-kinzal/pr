package pr

import (
	"context"

	"github.com/k-kinzal/pr/pkg/api"
)

type MergeOption struct {
	CommitTitleTemplate   string
	CommitMessageTemplate string
	MergeMethod           string
	*ListOption
}

func Merge(owner string, repo string, opt *MergeOption) ([]*api.PullRequest, error) {
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

	mergeOption := &api.MergeOption{
		CommitTitleTemplate:   opt.CommitTitleTemplate,
		CommitMessageTemplate: opt.CommitMessageTemplate,
		MergeMethod:           opt.MergeMethod,
	}
	mergedPulls, err := client.Merge(ctx, pulls, mergeOption)
	if err != nil {
		return nil, err
	}

	return mergedPulls, nil
}
