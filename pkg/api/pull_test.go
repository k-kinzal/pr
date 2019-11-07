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

func setup() {
	responder := func(callback func(length int) (interface{}, error)) httpmock.Responder {
		return func(req *http.Request) (*http.Response, error) {
			data, _ := callback(api.PerPage)
			jsonStr, _ := json.Marshal(data)
			resp := httpmock.NewStringResponse(200, string(jsonStr))
			resp.Header.Add("X-RateLimit-Limit", "5000")
			resp.Header.Add("X-RateLimit-Remaining", "4999")
			resp.Header.Add("X-RateLimit-Reset", fmt.Sprint(time.Now().Unix()))
			resp.Header.Add("Link", gen.LinkHeader(req.URL, 3))
			resp.Request = req
			return resp, nil
		}
	}
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/octocat/Hello-World/pulls", responder(func(length int) (interface{}, error) { return gen.PullRequests(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/comments`, responder(func(length int) (interface{}, error) { return gen.PullRequestComments(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/reviews`, responder(func(length int) (interface{}, error) { return gen.PullRequestReviews(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/pulls/\d+/commits`, responder(func(length int) (interface{}, error) { return gen.RepositoryCommits(length) }))
	httpmock.RegisterResponder("GET", `=~^https://api.github.com/repos/octocat/Hello-World/commits/[a-z0-9]+/statuses`, responder(func(length int) (interface{}, error) { return gen.RepoStatuses(length) }))
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func TestClient_GetPulls(t *testing.T) {
	setup()
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
}

func TestClient_GetPullsWithComments(t *testing.T) {
	setup()
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
}

func TestClient_GetPullsWithReviews(t *testing.T) {
	setup()
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
}

func TestClient_GetPullsWithCommits(t *testing.T) {
	setup()
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
}

func TestClient_GetPullsWithStatuses(t *testing.T) {
	setup()
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
}

func TestClient_GetPullsWithAll(t *testing.T) {
	setup()
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
}
