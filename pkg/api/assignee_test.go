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

func TestClient_AddAssignees(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees", func(request *http.Request) (response *http.Response, e error) {
		var req struct {
			Assignees []string `json:"assignees,omitempty"`
		}
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		issue, err := gen.Issue()
		if err != nil {
			return nil, err
		}
		issue.Assignees = make([]*github.User, len(issue.Assignees))
		for i, name := range req.Assignees {
			issue.Assignees[i] = &github.User{Login: &name}
		}

		resp, err := httpmock.NewJsonResponse(200, issue)
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
	opt := &api.AssigneeOption{
		Assignees: []string{"user1"},
	}
	pulls, err := client.AddAssignees(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Assignees) != 1 {
			t.Fatalf("len(pulls[%d].assignees): expect `1`, but actual `%d`", i, len(pull.Assignees))
		}
		if pull.Assignees[0].Login != "user1" {
			t.Fatalf("pulls[%d].assignees[0].name: expect `user1`, but actual `%s`", i, pull.Assignees[0].Login)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"], info)
	}
}

func TestClient_RemoveAssignees(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees", func(request *http.Request) (response *http.Response, e error) {
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
			Assignees: []*api.User{
				{
					Login: "user1",
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
			Assignees: []*api.User{
				{
					Login: "user1",
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
			Assignees: []*api.User{
				{
					Login: "user1",
				},
			},
		},
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.AssigneeOption{
		Assignees: []string{"user1"},
	}
	pulls, err := client.RemoveAssignees(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Assignees) != 0 {
			t.Fatalf("len(pulls[%d].labels): expect `0`, but actual `%d`", i, len(pull.Labels))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"], info)
	}
}

func TestClient_ReplaceAssignees(t *testing.T) {
	gen.Reset()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees", func(request *http.Request) (response *http.Response, e error) {
		var req struct {
			Assignees []string `json:"assignees,omitempty"`
		}
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}

		issue, err := gen.Issue()
		if err != nil {
			return nil, err
		}
		issue.Assignees = make([]*github.User, len(issue.Assignees))
		for i, name := range req.Assignees {
			issue.Assignees[i] = &github.User{Login: &name}
		}

		resp, err := httpmock.NewJsonResponse(200, issue)
		if err != nil {
			return nil, err
		}
		resp.Header.Add("X-RateLimit-Limit", "5000")
		resp.Header.Add("X-RateLimit-Remaining", "4999")
		resp.Header.Add("X-RateLimit-Reset", fmt.Sprint(time.Now().Unix()))
		resp.Request = request

		return resp, nil
	})
	httpmock.RegisterResponder("DELETE", "=~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees", func(request *http.Request) (response *http.Response, e error) {
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
			Assignees: []*api.User{
				{
					Login: "user1",
				},
				{
					Login: "user2",
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
			Assignees: []*api.User{
				{
					Login: "user1",
				},
				{
					Login: "user2",
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
			Assignees: []*api.User{
				{
					Login: "user1",
				},
				{
					Login: "user2",
				},
			},
		},
	}

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := &api.AssigneeOption{
		Assignees: []string{"user3"},
	}
	pulls, err := client.ReplaceAssignees(ctx, pulls, opt)
	if err != nil {
		t.Fatal(err)
	}
	for i, pull := range pulls {
		if len(pull.Assignees) != 1 {
			t.Fatalf("len(pulls[%d].labels): expect `1`, but actual `%d`", i, len(pull.Assignees))
		}
		if pull.Assignees[0].Login != "user3" {
			t.Fatalf("pulls[%d].labels[0].name: expect `user3`, but actual `%s`", i, pull.Assignees[0].Login)
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["POST =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"], info)
	}
	if info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["DELETE =~^https://api.github.com/repos/octocat/Hello-World/issues/\\d+/assignees"], info)
	}
}
