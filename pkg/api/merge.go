package api

import (
	"bytes"
	"context"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/google/go-github/v28/github"
)

const mergeRetryCount = 30
const mergeRetryIntervalSec = 2

type MergeOption struct {
	CommitTitleTemplate   string
	CommitMessageTemplate string
	MergeMethod           string
}

func (c *Client) Merge(ctx context.Context, pulls []*PullRequest, opt *MergeOption) ([]*PullRequest, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, pull := range pulls {
		eg.Go(func(pull *PullRequest) func() error {
			buf := &bytes.Buffer{}
			commitTitleTemplate := template.Must(template.New("commit title").Parse(opt.CommitTitleTemplate))
			commitMessageTemplate := template.Must(template.New("commit message").Parse(opt.CommitMessageTemplate))
			return func() error {
				if err := commitTitleTemplate.Execute(buf, pull); err != nil {
					return err
				}
				commitTitle := buf.String()
				buf.Reset()

				if err := commitMessageTemplate.Execute(buf, pull); err != nil {
					return err
				}
				commitMessage := buf.String()
				buf.Reset()

				state := "closed"
				now := Timestamp(time.Now().UTC().Unix()) //.Format("2006-01-02 15:04:05 -0700")

				o := &github.PullRequestOptions{
					CommitTitle: commitTitle,
					MergeMethod: opt.MergeMethod,
				}
				for i := 0; i < mergeRetryCount; i++ {
					result, _, err := c.github.PullRequests.Merge(ctx, pull.Owner, pull.Repo, int(pull.Number), commitMessage, o)
					if err != nil {
						if e, ok := err.(*github.ErrorResponse); ok && e.Response.StatusCode == http.StatusMethodNotAllowed {
							if i+1 == mergeRetryCount {
								return err
							}
							time.Sleep(time.Second * mergeRetryIntervalSec)
							continue
						}
						return err
					}
					if result.GetMerged() {
						pull.State = state
						pull.UpdatedAt = now
						pull.ClosedAt = now
						pull.MergedAt = now
						pull.MergeCommitSha = *result.SHA
						break
					}
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
