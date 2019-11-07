package pr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/xerrors"

	"github.com/k-kinzal/pr/pkg/api"
)

type MergeOption struct {
	CommitTitleTemplate   string
	CommitMessageTemplate string
	MergeMethod           string
	PullOption
}

func Merge(opt MergeOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOption := &api.Options{
		Token:     opt.Token,
		RateLimit: opt.Rate,
	}
	client := api.NewClient(ctx, clientOption)

	pullOption := api.PullsOption{
		EnableComments: opt.EnableComments,
		EnableReviews:  opt.EnableReviews,
		EnableCommits:  opt.EnableCommits,
		EnableStatuses: opt.EnableStatuses,
		Rules:          api.NewPullRequestRules(opt.Rules, opt.Limit),
	}
	pulls, err := client.GetPulls(ctx, opt.Owner, opt.Repo, pullOption)
	if err != nil {
		return err
	}
	if len(pulls) == 0 {
		fmt.Fprintln(os.Stdout, "[]")
		return &NoMatchError{pullOption.Rules}
	}

	mergeOption := api.MergeOption{
		CommitTitleTemplate:   opt.CommitTitleTemplate,
		CommitMessageTemplate: opt.CommitMessageTemplate,
		MergeMethod:           opt.MergeMethod,
	}
	mergedPulls, err := client.Merge(ctx, pulls, mergeOption)
	if err != nil {
		return xerrors.Errorf("merge: %s", err)
	}

	out, err := json.Marshal(mergedPulls)
	if err != nil {
		return xerrors.Errorf("merge: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
