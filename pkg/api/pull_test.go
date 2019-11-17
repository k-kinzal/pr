package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/k-kinzal/pr/test/gen"

	"github.com/jarcoal/httpmock"
	"github.com/k-kinzal/pr/pkg/api"
)

func setup(page int) {
	responder := func(callback func(length int) (interface{}, error)) httpmock.Responder {
		return func(req *http.Request) (*http.Response, error) {
			data, _ := callback(api.PerPage)
			jsonStr, _ := json.Marshal(data)
			resp := httpmock.NewStringResponse(200, string(jsonStr))
			resp.Header.Add("X-RateLimit-Limit", "5000")
			resp.Header.Add("X-RateLimit-Remaining", "4999")
			resp.Header.Add("X-RateLimit-Reset", fmt.Sprint(time.Now().Unix()))
			resp.Header.Add("Link", gen.LinkHeader(req.URL, page))
			resp.Request = req
			return resp, nil
		}
	}
	gen.Reset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/octocat/Hello-World/pulls/1", responder(func(length int) (i interface{}, e error) { return gen.PullRequest() }))
	httpmock.RegisterResponder("GET", "=~^https://api.github.com/search/issues", responder(func(length int) (interface{}, error) { return gen.IssuesSearchResult() }))
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/octocat/Hello-World/pulls", responder(func(length int) (interface{}, error) { return gen.PullRequests(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`, responder(func(length int) (interface{}, error) { return gen.PullRequestComments(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`, responder(func(length int) (interface{}, error) { return gen.PullRequestReviews(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`, responder(func(length int) (interface{}, error) { return gen.RepositoryCommits(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`, responder(func(length int) (interface{}, error) { return gen.RepoStatuses(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`, responder(func(length int) (interface{}, error) { return gen.ListCheckRunsResults() }))
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func TestClient_GetPulls(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: false,
		EnableReviews:  false,
		EnableCommits:  false,
		EnableStatuses: false,
		EnableChecks:   false,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 0 {
			t.Fatalf("pulls[%d]: expect `%d`, but actual `%d`", i, 0, len(pull.Comments))
		}
		if len(pull.Reviews) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithComments(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: true,
		EnableReviews:  false,
		EnableCommits:  false,
		EnableStatuses: false,
		EnableChecks:   false,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithReviews(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: false,
		EnableReviews:  true,
		EnableCommits:  false,
		EnableStatuses: false,
		EnableChecks:   false,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 300 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 0 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithCommits(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: false,
		EnableReviews:  false,
		EnableCommits:  true,
		EnableStatuses: false,
		EnableChecks:   false,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithStatuses(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: false,
		EnableReviews:  false,
		EnableCommits:  false,
		EnableStatuses: true,
		EnableChecks:   false,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithChecks(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: false,
		EnableReviews:  false,
		EnableCommits:  false,
		EnableStatuses: false,
		EnableChecks:   true,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 0 {
			t.Fatalf("pulls[%d]: expect `0`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 3 {
			t.Fatalf("pulls[%d]: expect `3`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsWithAll(t *testing.T) {
	setup(3)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: true,
		EnableReviews:  true,
		EnableCommits:  true,
		EnableStatuses: true,
		EnableChecks:   true,
		Rules:          api.NewPullRequestRules([]string{}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != opt.Rules.GetLimit() {
		t.Fatalf("pulls: expect `%d`, but actual `%d`", opt.Rules.GetLimit(), len(pulls))
	}
	for i, pull := range pulls {
		if len(pull.Comments) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Comments))
		}
		if len(pull.Reviews) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Reviews))
		}
		if len(pull.Commits) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Commits))
		}
		if len(pull.Statuses) != 300 {
			t.Fatalf("pulls[%d]: expect `300`, but actual `%d`", i, len(pull.Statuses))
		}
		if len(pull.Checks) != 3 {
			t.Fatalf("pulls[%d]: expect `3`, but actual `%d`", i, len(pull.Checks))
		}
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 3 {
		t.Fatalf("expect `3`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 900 {
		t.Fatalf("expect `900`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsPassNumber(t *testing.T) {
	setup(1)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: true,
		EnableReviews:  true,
		EnableCommits:  true,
		EnableStatuses: true,
		EnableChecks:   true,
		Rules: api.NewPullRequestRules([]string{
			"number == `1`",
		}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != 1 {
		t.Fatalf("pulls: expect `1`, but actual `%d`", len(pulls))
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls/1"] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls/1"], info)
	}
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}

func TestClient_GetPullsPassSha(t *testing.T) {
	setup(1)
	defer teardown()

	ctx := context.Background()
	client := api.NewClient(ctx, &api.Options{
		Token:     "xxxx",
		RateLimit: math.MaxInt32,
	})
	opt := api.PullsOption{
		EnableComments: true,
		EnableReviews:  true,
		EnableCommits:  true,
		EnableStatuses: true,
		EnableChecks:   true,
		Rules: api.NewPullRequestRules([]string{
			"head.sha == `\"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b\"`",
		}, 300),
	}
	pulls, err := client.GetPulls(ctx, "octocat", "Hello-World", opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(pulls) != 1 {
		t.Fatalf("pulls: expect `1`, but actual `%d`", len(pulls))
	}

	info := httpmock.GetCallCountInfo()
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls/1"] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls/1"], info)
	}
	if info["GET =~^https://api.github.com/search/issues"] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info["GET =~^https://api.github.com/search/issues"], info)
	}
	if info["GET https://api.github.com/repos/octocat/Hello-World/pulls"] != 0 {
		t.Fatalf("expect `0`, but actual `%d`: %#v", info["GET https://api.github.com/repos/octocat/Hello-World/pulls"], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`], info)
	}
	if info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`] != 1 {
		t.Fatalf("expect `1`, but actual `%d`: %#v", info[`GET =~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/check-runs`], info)
	}
}
