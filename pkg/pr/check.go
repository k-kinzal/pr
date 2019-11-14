package pr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/action"

	"github.com/k-kinzal/pr/pkg/api"
	"golang.org/x/xerrors"
)

type CheckOption struct {
	TargetURL      string
	Limit          int
	Rate           int
	Rules          []string
	EnableComments bool
	EnableReviews  bool
	EnableCommits  bool
	EnableStatuses bool
	Action         string
	MergeOption    MergeOption
	Option
}

func Check(opt CheckOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOption := &api.Options{
		Token:     opt.Token,
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
	pulls, err := client.GetPulls(ctx, opt.Owner, opt.Repo, pullOption)
	if err != nil {
		return xerrors.Errorf("check: %s", err)
	}

	if opt.TargetURL == "" && action.Actions {
		opt.TargetURL = "https://github.com/{{ .Owner }}/{{ .Repo }}/commit/{{ .Head.Sha }}/checks"
	}
	checkOption := &api.CheckOption{
		State:       "pending",
		TargetURL:   opt.TargetURL,
		Description: "Checking for matching rules",
		Context:     "PR",
	}
	checked, err := client.Check(ctx, pulls, checkOption)
	if err != nil {
		return xerrors.Errorf("check: %s", err)
	}

	filterd, filterdErr := rules.Apply(checked)
	if filterdErr != nil {
		return xerrors.Errorf("check: %s", err)
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

	checkOption.State = "success"
	checkOption.Description = "Matched the rules"
	checkedMatches, err := client.Check(ctx, matches, checkOption)
	if err != nil {
		return xerrors.Errorf("check: %s", err)
	}
	checkOption.State = "pending"
	checkOption.Description = fmt.Sprintf("Does not match `%s`", rules.Expression())
	checkeNomatches, err := client.Check(ctx, nomatches, checkOption)
	if err != nil {
		return xerrors.Errorf("check: %s", err)
	}

	var actioned []*api.PullRequest
	switch opt.Action {
	case "merge":
		mergeOption := api.MergeOption{
			CommitTitleTemplate:   opt.MergeOption.CommitTitleTemplate,
			CommitMessageTemplate: opt.MergeOption.CommitMessageTemplate,
			MergeMethod:           opt.MergeOption.MergeMethod,
		}
		merged, err := client.Merge(ctx, checkedMatches, mergeOption)
		if err != nil {
			return xerrors.Errorf("check: %s", err)
		}
		actioned = merged
	default:
		actioned = checkedMatches
	}

	out, err := json.Marshal(append(actioned, checkeNomatches...))
	if err != nil {
		return xerrors.Errorf("check: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
