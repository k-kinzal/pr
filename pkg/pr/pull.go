package pr

import (
	"context"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/api"
)

type NoMatchError struct {
	rules *api.PullRequestRules
}

func (e *NoMatchError) Error() string {
	return fmt.Sprintf("no PR matches the rule: `%s`", e.rules.Expression())
}

type ListOption struct {
	Limit          int
	Rate           int
	Rules          []string
	EnableComments bool
	EnableReviews  bool
	EnableCommits  bool
	EnableStatuses bool
	EnableChecks   bool
}

func List(owner string, repo string, opt *ListOption) ([]*api.PullRequest, error) {
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
		fmt.Fprintln(os.Stdout, "[]")
		return nil, &NoMatchError{pullOption.Rules}
	}

	return pulls, nil
}
