package pr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/api"
)

type NoMatchError struct {
	filter *api.PullsFilter
}

func (e *NoMatchError) Error() string {
	return fmt.Sprintf("no PR matches the rule: `%s`", e.filter.Expression())
}

type Option struct {
	Token string
	Owner string
	Repo  string
}

type PullOption struct {
	*Option
	Limit int
	Rules []string
}

func Show(opt *PullOption) error {
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

	out, err := json.Marshal(pulls)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
