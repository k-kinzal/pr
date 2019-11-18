package pr

import (
	"context"
	"errors"
	"fmt"

	"github.com/k-kinzal/pr/pkg/api"
)

type ValidationResult interface {
	Success() bool
	String() string
}

type ValidationResultSuccess struct {
	rule    string
	message string
}

func (v *ValidationResultSuccess) Success() bool {
	return true
}

func (v *ValidationResultSuccess) String() string {
	return fmt.Sprintf("%s: %s", v.rule, v.message)
}

type ValidationResultFailed struct {
	rule string
	err  error
}

func (v *ValidationResultFailed) Success() bool {
	return false
}

func (v *ValidationResultFailed) String() string {
	return fmt.Sprintf("%s: %s", v.rule, v.err)
}

func Validate(owner string, repo string, opt *ListOption) ([]ValidationResult, []*api.PullRequest) {
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
		Rules:          api.NewPullRequestRules([]string{opt.Rules[0]}, opt.Limit),
	}
	pulls, err := client.GetPulls(ctx, owner, repo, pullOption)
	if err != nil {
		result := []ValidationResult{&ValidationResultFailed{opt.Rules[0], err}}
		return result, nil
	}

	r := api.NewPullRequestRules([]string{}, opt.Limit)
	var results []ValidationResult
	for _, rule := range opt.Rules {
		r.Add(rule)
		p, err := r.Apply(pulls)
		pulls = p
		if err != nil {
			results = append(results, &ValidationResultFailed{rule, err})
			continue
		}
		if len(pulls) == 0 {
			err := errors.New("no PR matches the rule")
			results = append(results, &ValidationResultFailed{rule, err})
			continue
		}
		message := fmt.Sprintf("%d PRs matched the rules", len(pulls))
		results = append(results, &ValidationResultSuccess{rule, message})
	}

	return results, pulls
}
