package pr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/api"
)

type MergeOption struct {
	CommitTitleTemplate   string
	CommitMessageTemplate string
	MergeMethod           string
	*PullOption
}

func Merge(opt *MergeOption) error {
	client := api.NewClient(opt.Token)
	filter := api.NewPullsFilter(opt.Rules, opt.Limit)
	pulls, err := client.GetPulls(context.Background(), opt.Owner, opt.Repo, filter)
	if err != nil {
		return err
	}
	if len(pulls) == 0 {
		fmt.Fprintln(os.Stdout, "[]")
		return &NoMatchError{filter}
	}

	o := &api.MergeOption{
		CommitTitleTemplate:   opt.CommitTitleTemplate,
		CommitMessageTemplate: opt.CommitMessageTemplate,
		MergeMethod:           opt.MergeMethod,
	}
	mergedPulls, err := client.Merge(context.Background(), pulls, o)
	if err != nil {
		return err
	}

	out, err := json.Marshal(mergedPulls)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
