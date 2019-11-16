package api

import (
	"bytes"
	"context"
	"text/template"

	"golang.org/x/sync/errgroup"

	"github.com/google/go-github/v28/github"
)

type StatusOption struct {
	State       string
	TargetURL   string
	Description string
	Context     string
}

func (c *Client) Status(ctx context.Context, pulls []*PullRequest, opt *StatusOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			buf := &bytes.Buffer{}
			targetUrlTemplate := template.Must(template.New("target url").Parse(opt.TargetURL))
			descriptionTemplate := template.Must(template.New("description").Parse(opt.Description))
			contextTemplate := template.Must(template.New("context").Parse(opt.Context))
			return func() error {
				if err := targetUrlTemplate.Execute(buf, pull); err != nil {
					return err
				}
				targetUrl := buf.String()
				buf.Reset()

				if err := descriptionTemplate.Execute(buf, pull); err != nil {
					return err
				}
				description := buf.String()
				buf.Reset()

				if err := contextTemplate.Execute(buf, pull); err != nil {
					return err
				}
				context := buf.String()
				buf.Reset()

				status := &github.RepoStatus{
					State:       &opt.State,
					Description: &description,
					TargetURL:   &targetUrl,
					Context:     &context,
				}
				data, _, err := c.github.Repositories.CreateStatus(ctx, pull.Owner, pull.Repo, pull.Head.Sha, status)
				if err != nil {
					return err
				}
				var find bool
				for i, status := range pull.Statuses {
					if data.GetContext() == status.Context {
						find = true
						pull.Statuses[i] = newRepoStatus(data)
					}
				}
				if !find {
					pull.Statuses = append(pull.Statuses, newRepoStatus(data))
				}
				return nil
			}
		}(pull))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return pulls, nil
}
