package pr

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/k-kinzal/pr/pkg/api"
)

type AssigneeAction string

const (
	AssigneeActionAppend  AssigneeAction = "append"
	AssigneeActionRemove  AssigneeAction = "remove"
	AssigneeActionReplace AssigneeAction = "replace"
)

type AssigneeFunc func(assignees []string) []string

func RandomizeAssignee(assignees []string) []string {
	rand.Seed(time.Now().UnixNano())
	return []string{assignees[rand.Intn(len(assignees)-1)]}
}

type AssigneeOption struct {
	Assignees []string
	Action    AssigneeAction
	FuncList  []AssigneeFunc
	*ListOption
}

func Assignee(owner string, repo string, opt *AssigneeOption) ([]*api.PullRequest, error) {
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

	assignees := opt.Assignees
	for _, fn := range opt.FuncList {
		assignees = fn(assignees)
	}

	switch opt.Action {
	case AssigneeActionAppend:
		option := &api.AssigneeOption{
			Assignees: assignees,
		}
		return client.AddAssignees(ctx, pulls, option)
	case AssigneeActionRemove:
		option := &api.AssigneeOption{
			Assignees: assignees,
		}
		return client.RemoveAssignees(ctx, pulls, option)
	case AssigneeActionReplace:
		option := &api.AssigneeOption{
			Assignees: assignees,
		}
		return client.ReplaceAssignees(ctx, pulls, option)
	}

	return nil, fmt.Errorf("action is expected to be `append`, `remove`, `all`, but was actually %s", opt.Action)
}
