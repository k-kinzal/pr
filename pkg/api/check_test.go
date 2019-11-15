package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/k-kinzal/pr/pkg/api"
	"github.com/k-kinzal/pr/test/gen"
)

func TestClient_Check(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e", func(request *http.Request) (response *http.Response, e error) {
		status, err := gen.RepoStatus()
		if err != nil {
			return nil, err
		}
		if err := json.NewDecoder(request.Body).Decode(status); err != nil {
			return nil, err
		}
		resp, err := httpmock.NewJsonResponse(200, status)
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
	opt := &api.CheckOption{
		State:       "pending",
		TargetURL:   "https://github.com/{{ .Owner }}/{{ .Repo }}/commit/{{ .Head.Sha }}/checks",
		Description: "Test check",
		Context:     "PR",
	}
	pulls, err := client.Check(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Statuses) != 1 {
			t.Fatalf("len(pulls[%d].statuses): expect `%d`, but actual `%d`", i, 1, len(pull.Statuses))
		}
		if pull.Statuses[0].State != opt.State {
			t.Fatalf("pulls[%d].statuses[0].state): expect `%s`, but actual `%s`", i, opt.State, pull.Statuses[0].State)
		}
		targetUrl := fmt.Sprintf("https://github.com/%s/%s/commit/%s/checks", pull.Owner, pull.Repo, pull.Head.Sha)
		if pull.Statuses[0].TargetUrl != targetUrl {
			t.Fatalf("pulls[%d].statuses[0].target_url): expect `%s`, but actual `%s`", i, targetUrl, pull.Statuses[0].TargetUrl)
		}
		if pull.Statuses[0].Description != opt.Description {
			t.Fatalf("pulls[%d].statuses[0].description): expect `%s`, but actual `%s`", i, opt.Description, pull.Statuses[0].Description)
		}
		if pull.Statuses[0].Context != opt.Context {
			t.Fatalf("pulls[%d].statuses[0].context): expect `%s`, but actual `%s`", i, opt.Context, pull.Statuses[0].Context)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["POST https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["POST https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e"], info)
	}
}
