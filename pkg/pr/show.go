package pr

import (
	"context"
	"encoding/json"
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
	client := api.NewClient(opt.Token)
	pullOption := api.PullsOption{
		Rules:     api.NewPullRequestRules(opt.Rules, opt.Limit),
		RateLimit: opt.Rate,
	}
	pulls, err := client.GetPulls(context.Background(), opt.Owner, opt.Repo, pullOption)
	if err != nil {
		return err
	}
	if len(pulls) == 0 {
		fmt.Fprintln(os.Stdout, "[]")
		return &NoMatchError{pullOption.Rules}
	}

	out, err := json.Marshal(pulls)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(out))

	return nil
}
