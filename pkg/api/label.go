package api

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type LabelOption struct {
	Labels []string
}

func (c *Client) AppendLabel(ctx context.Context, pulls []*PullRequest, opt *LabelOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				now := Timestamp(time.Now().UTC().Unix())
				labels, _, err := c.github.Issues.AddLabelsToIssue(ctx, pull.Owner, pull.Repo, int(pull.Number), opt.Labels)
				if err != nil {
					return err
				}

				for _, label := range labels {
					pull.Labels = append(pull.Labels, newLabel(label))
				}
				pull.UpdatedAt = now

				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}

func (c *Client) RemoveLabel(ctx context.Context, pulls []*PullRequest, opt *LabelOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		for _, label := range opt.Labels {
			eg.Go(func(pull *PullRequest, label string) func() error {
				return func() error {
					now := Timestamp(time.Now().UTC().Unix())
					_, err := c.github.Issues.RemoveLabelForIssue(ctx, pull.Owner, pull.Repo, int(pull.Number), label)
					if err != nil {
						return err
					}
					for i, lbl := range pull.Labels {
						if lbl.Name == label {
							pull.Labels[i] = pull.Labels[len(pull.Labels)-1]
							pull.Labels = pull.Labels[:len(pull.Labels)-1]
							break
						}
					}
					pull.UpdatedAt = now
					return nil
				}
			}(pull, label))
		}
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}

func (c *Client) ReplaceLabel(ctx context.Context, pulls []*PullRequest, opt *LabelOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			return func() error {
				now := Timestamp(time.Now().UTC().Unix())
				labels, _, err := c.github.Issues.ReplaceLabelsForIssue(ctx, pull.Owner, pull.Repo, int(pull.Number), opt.Labels)
				if err != nil {
					return err
				}
				pull.Labels = make([]*Label, len(labels))
				for i, label := range labels {
					pull.Labels[i] = newLabel(label)
				}
				pull.UpdatedAt = now

				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pulls, nil
}
