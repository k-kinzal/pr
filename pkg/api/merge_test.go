package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-github/v28/github"

	"github.com/jarcoal/httpmock"
	"github.com/k-kinzal/pr/pkg/api"
	"github.com/k-kinzal/pr/test/gen"
)

func TestClient_Merge(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "=~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/merge", func(request *http.Request) (response *http.Response, e error) {
		var req struct {
			CommitMessage string `json:"commit_message"`
			CommitTitle   string `json:"commit_title,omitempty"`
			MergeMethod   string `json:"merge_method,omitempty"`
			SHA           string `json:"sha,omitempty"`
		}
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		result := github.PullRequestMergeResult{
			SHA:     &req.SHA,
			Merged:  func(b bool) *bool { return &b }(true),
			Message: &req.CommitMessage,
		}

		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			return nil, err
		}
		resp.Header.Add("X-RateLimit-Limit", "5000")
		resp.Header.Add("X-RateLimit-Remaining", "4999")
		resp.Header.Add("X-RateLimit-Reset", fmt.Sprint(time.Now().Unix()))
		resp.Request = request

		return resp, nil
	})

	pulls := []*api.PullRequest{
		{
			Id:     1,
			Number: 1,
			State:  "open",
			Head: &api.PullRequestBranch{
				Sha: "6dcb09b5b57875f334f61aebed695e2e4193db5e",
			},
			Owner: "octocat",
			Repo:  "Hello-World",
		},
		{
			Id:     2,
			Number: 2,
			State:  "open",
			Head: &api.PullRequestBranch{
				Sha: "6dcb09b5b57875f334f61aebed695e2e4193db5e",
			},
			Owner: "octocat",
			Repo:  "Hello-World",
		},
		{
			Id:     3,
			Number: 3,
			State:  "open",
			Head: &api.PullRequestBranch{
				Sha: "6dcb09b5b57875f334f61aebed695e2e4193db5e",
			},
			Owner: "octocat",
			Repo:  "Hello-World",
		},
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.MergeOption{
		CommitTitleTemplate:   "Merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}",
		CommitMessageTemplate: "{{ .Title }}",
	}
	pulls, err := client.Merge(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if pull.State != "closed" {
			t.Fatalf("pulls[%d].state: expect `closed`, but actual `%s`", i, pull.Statuses[0].State)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["PUT =~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/merge"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["PUT =~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/merge"], info)
	}
}
