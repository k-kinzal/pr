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

func TestClient_AppendLabel(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels", func(request *http.Request) (response *http.Response, e error) {
		var req []string
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		labels, err := gen.Labels(len(req))
		if err != nil {
			return nil, err
		}
		for i, name := range req {
			labels[i].Name = &name
		}

		resp, err := httpmock.NewJsonResponse(200, labels)
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
	opt := &api.LabelOption{
		Labels: []string{"label1"},
	}
	pulls, err := client.AppendLabel(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Labels) != 1 {
			t.Fatalf("len(pulls[%d].labels): expect `1`, but actual `%d`", i, len(pull.Labels))
		}
		if pull.Labels[0].Name != "label1" {
			t.Fatalf("pulls[%d].labels[0].name: expect `label1`, but actual `%s`", i, pull.Labels[0].Name)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels"], info)
	}
}

func TestClient_RemoveLabel(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels/.+", func(request *http.Request) (response *http.Response, e error) {
		resp, err := httpmock.NewJsonResponse(200, nil)
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
		},
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.LabelOption{
		Labels: []string{"label1"},
	}
	pulls, err := client.RemoveLabel(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Labels) != 0 {
			t.Fatalf("len(pulls[%d].labels): expect `0`, but actual `%d`", i, len(pull.Labels))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels/.+"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels/.+"], info)
	}
}

func TestClient_ReplaceLabel(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels", func(request *http.Request) (response *http.Response, e error) {
		var req []string
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		labels, err := gen.Labels(len(req))
		if err != nil {
			return nil, err
		}
		for i, name := range req {
			labels[i].Name = &name
		}
		resp, err := httpmock.NewJsonResponse(200, labels)
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
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
			Labels: []*api.Label{
				{
					Name: "label1",
				},
			},
		},
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.LabelOption{
		Labels: []string{"label2"},
	}
	pulls, err := client.ReplaceLabel(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Labels) != 1 {
			t.Fatalf("len(pulls[%d].labels): expect `1`, but actual `%d`", i, len(pull.Labels))
		}
		if pull.Labels[0].Name != "label2" {
			t.Fatalf("pulls[%d].labels[0].name: expect `label2`, but actual `%s`", i, pull.Labels[0].Name)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["PUT =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["PUT =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/labels"], info)
	}
}
