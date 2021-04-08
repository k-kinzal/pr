package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v28/github"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/k-kinzal/pr/pkg/api"
	"github.com/k-kinzal/pr/test/gen"
)

func TestClient_AddApproval(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "=~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/reviews", func(request *http.Request) (response *http.Response, e error) {
		var req struct {
			Action *string `json:"event"`
		}
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		pullRequestReview, err := gen.PullRequestReview()
		if err != nil {
			return nil, err
		}

		pullRequestReview.State = github.String("APPROVED")

		resp, err := httpmock.NewJsonResponse(200, pullRequestReview)
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
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.ReviewOption{
		Action: "approve",
	}

	pulls, err := client.AddApproval(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		for j, review := range pull.Reviews {
			if review.State != "APPROVED" {
				t.Fatalf("pull[`%d`].reviews[`%d`]): expect APPROVED, but actual `%s`", i, j, review.State)
			}
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["POST =~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/reviews"] != 2 {
		t.Fatalf("expect `2`, but actual `%d`: %#v", info["POST =~^https://api.github.com/repos/octocat/Hello-World/pulls/\\d+/reviews"], info)
	}
}
