package pr

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/k-kinzal/pr/pkg/api"
)

type LabelAction string

const (
	LabelActionAppend  LabelAction = "append"
	LabelActionRemove  LabelAction = "remove"
	LabelActionReplace LabelAction = "replace"
)

type LabelFunc func(labels []string) []string

func RandomizeLabel(labels []string) []string {
	rand.Seed(time.Now().UnixNano())
	// #nosec
	return []string{labels[rand.Intn(len(labels)-1)]}
}

type LabelOption struct {
	Labels   []string
	Action   LabelAction
	FuncList []LabelFunc
	*ListOption
}

func Label(owner string, repo string, opt *LabelOption) ([]*api.PullRequest, error) {
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

	labels := opt.Labels
	for _, fn := range opt.FuncList {
		labels = fn(labels)
	}

	switch opt.Action {
	case LabelActionAppend:
		labelOption := &api.LabelOption{
			Labels: labels,
		}
		return client.AppendLabel(ctx, pulls, labelOption)
	case LabelActionRemove:
		labelOption := &api.LabelOption{
			Labels: labels,
		}
		return client.RemoveLabel(ctx, pulls, labelOption)
	case LabelActionReplace:
		labelOption := &api.LabelOption{
			Labels: labels,
		}
		return client.ReplaceLabel(ctx, pulls, labelOption)
	}

	return nil, fmt.Errorf("action is expected to be `append`, `remove`, `all`, but was actually %s", opt.Action)
}
