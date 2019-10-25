package pr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/api"
	"golang.org/x/xerrors"
)

type NoMatchError struct {
	rules *api.PullRequestRules
}

func (e *NoMatchError) Error() string {
	return fmt.Sprintf("no PR matches the rule: `%s`", e.rules.Expression())
}

type Option struct {
	Token string
	Owner string
	Repo  string
}

type PullOption struct {
	Limit int
	Rate  int
	Rules []string
	Option
}

func Show(opt PullOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOption := &api.Options{
		Token:     opt.Token,
		RateLimit: opt.Rate,
	}
	client := api.NewClient(ctx, clientOption)

	pullOption := api.PullsOption{
		Rules: api.NewPullRequestRules(opt.Rules, opt.Limit),
	}
	pulls, err := client.GetPulls(ctx, opt.Owner, opt.Repo, pullOption)
	if err != nil {
		return xerrors.Errorf("show: %s", err)
	}
	if len(pulls) == 0 {
		fmt.Fprintln(os.Stdout, "[]")
		return &NoMatchError{pullOption.Rules}
	}

	out, err := json.Marshal(pulls)
	if err != nil {
		return xerrors.Errorf("show: %s", err)
	}

	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
