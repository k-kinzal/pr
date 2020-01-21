package api

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type AssigneeOption struct {
	Assignees []string
}

func (c *Client) AddAssignees(ctx context.Context, pulls []*PullRequest, opt *AssigneeOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				issue, _, err := c.github.Issues.AddAssignees(ctx, pull.Owner, pull.Repo, int(pull.Number), opt.Assignees)
				if err != nil {
					return err
				}

				users := make([]*User, len(issue.Assignees))
				for i, assignee := range issue.Assignees {
					users[i] = newUser(assignee)
				}
				pull.Assignees = users
				pull.UpdatedAt = newTimestamp(issue.UpdatedAt)

				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}

func (c *Client) RemoveAssignees(ctx context.Context, pulls []*PullRequest, opt *AssigneeOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				issue, _, err := c.github.Issues.RemoveAssignees(ctx, pull.Owner, pull.Repo, int(pull.Number), opt.Assignees)
				if err != nil {
					return err
				}

				users := make([]*User, len(issue.Assignees))
				for i, assignee := range issue.Assignees {
					users[i] = newUser(assignee)
				}
				pull.Assignees = users
				pull.UpdatedAt = newTimestamp(issue.UpdatedAt)

				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}

func (c *Client) ReplaceAssignees(ctx context.Context, pulls []*PullRequest, opt *AssigneeOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				assignees := make([]string, len(pull.Assignees))
				for i, a := range pull.Assignees {
					assignees[i] = a.Login
				}
				_, _, err := c.github.Issues.RemoveAssignees(ctx, pull.Owner, pull.Repo, int(pull.Number), assignees)
				if err != nil {
					return err
				}
				issue, resp, err := c.github.Issues.AddAssignees(ctx, pull.Owner, pull.Repo, int(pull.Number), opt.Assignees)
				if err != nil {
					return err
				}
				fmt.Println(resp)

				users := make([]*User, len(issue.Assignees))
				for i, assignee := range issue.Assignees {
					users[i] = newUser(assignee)
				}
				pull.Assignees = users
				pull.UpdatedAt = newTimestamp(issue.UpdatedAt)

				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}
