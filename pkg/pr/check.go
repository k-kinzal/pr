package pr

import (
	"context"
	"fmt"

	"github.com/k-kinzal/pr/pkg/action"

	"github.com/k-kinzal/pr/pkg/api"
)

type CheckOption struct {
	TargetURL string
	Action    string
	*MergeOption
	*ListOption
}

func Check(owner string, repo string, opt *CheckOption) ([]*api.PullRequest, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOption := &api.Options{
		Token:     token,
		RateLimit: opt.Rate,
	}
	client := api.NewClient(ctx, clientOption)

	rules := api.NewPullRequestRules(opt.Rules, opt.Limit)
	if state := rules.GetState(); state != "open" {
		rules.Add("state == `\"open\"`")
	}

	pullOption := api.PullsOption{
		EnableComments: opt.EnableComments,
		EnableReviews:  opt.EnableReviews,
		EnableCommits:  opt.EnableCommits,
		EnableStatuses: true,
		Rules:          rules.SearchRules(),
	}
	pulls, err := client.GetPulls(ctx, owner, repo, pullOption)
	if err != nil {
		return nil, err
	}

	if opt.TargetURL == "" && action.Actions {
		opt.TargetURL = "https://github.com/{{ .Owner }}/{{ .Repo }}/commit/{{ .Head.Sha }}/checks"
	}
	statusOption := &api.StatusOption{
		State:       "pending",
		TargetURL:   opt.TargetURL,
		Description: "Checking for matching rules",
		Context:     "PR",
	}
	checked, err := client.Status(ctx, pulls, statusOption)
	if err != nil {
		return nil, err
	}

	filterd, filterdErr := rules.Apply(checked)
	if filterdErr != nil {
		return nil, err
	}
	matches := make([]*api.PullRequest, 0)
	nomatches := make([]*api.PullRequest, 0)
	for _, pull1 := range pulls {
		var match bool
		for _, pull2 := range filterd {
			if match = pull1.Id == pull2.Id; match {
				matches = append(matches, pull1)
				break
			}
		}
		if !match {
			nomatches = append(nomatches, pull1)
		}
	}

	statusOption.State = "success"
	statusOption.Description = "Matched the rules"
	checkedMatches, err := client.Status(ctx, matches, statusOption)
	if err != nil {
		return nil, err
	}
	statusOption.State = "pending"
	statusOption.Description = fmt.Sprintf("Does not match `%s`", rules.Expression())
	checkeNomatches, err := client.Status(ctx, nomatches, statusOption)
	if err != nil {
		return nil, err
	}

	var actioned []*api.PullRequest
	switch opt.Action {
	case "merge":
		mergeOption := &api.MergeOption{
			CommitTitleTemplate:   opt.MergeOption.CommitTitleTemplate,
			CommitMessageTemplate: opt.MergeOption.CommitMessageTemplate,
			MergeMethod:           opt.MergeOption.MergeMethod,
		}
		merged, err := client.Merge(ctx, checkedMatches, mergeOption)
		if err != nil {
			return nil, err
		}
		actioned = merged
	default:
		actioned = checkedMatches
	}

	return append(actioned, checkeNomatches...), nil
}
