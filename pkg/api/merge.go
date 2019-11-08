package api

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"golang.org/x/xerrors"

	"github.com/google/go-github/v28/github"
)

type MergeOption struct {
	CommitTitleTemplate   string
	CommitMessageTemplate string
	MergeMethod           string
}

func (c *Client) Merge(ctx context.Context, pulls []*PullRequest, opt MergeOption) ([]*PullRequest, error) {
	buf := &bytes.Buffer{}
	commitTitleTemplate := template.Must(template.New("commit title").Parse(opt.CommitTitleTemplate))
	commitMessageTemplate := template.Must(template.New("commit message").Parse(opt.CommitMessageTemplate))

	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan *PullRequest)
	errCh := make(chan error)
	for _, pull := range pulls {
		go func(pull *PullRequest) {
			if err := commitTitleTemplate.Execute(buf, pull); err != nil {
				errCh <- err
				return
			}
			commitTitle := buf.String()
			buf.Reset()

			if err := commitMessageTemplate.Execute(buf, pull); err != nil {
				errCh <- err
				return
			}
			commitMessage := buf.String()
			buf.Reset()

			state := "closed"
			now := Timestamp(time.Now().UTC().Unix()) //.Format("2006-01-02 15:04:05 -0700")

			o := &github.PullRequestOptions{
				CommitTitle: commitTitle,
				MergeMethod: opt.MergeMethod,
			}
			result, _, err := c.github.PullRequests.Merge(childCtx, pull.Owner, pull.Repo, int(pull.Number), commitMessage, o)
			if err != nil {
				errCh <- err
				return
			}
			if !result.GetMerged() {
				errCh <- xerrors.New(result.GetMessage())

			}

			p := *pull
			p.State = state
			p.UpdatedAt = now
			p.ClosedAt = now
			p.MergedAt = now
			p.MergeCommitSha = *result.SHA

			ch <- &p
		}(pull)
	}

	var allPulls []*PullRequest
	for i := 0; i < len(pulls); i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			break
		case pull := <-ch:
			allPulls = append(allPulls, pull)
		case err := <-errCh:
			cancel()
			return nil, err
		}
	}
	return allPulls, nil
}
