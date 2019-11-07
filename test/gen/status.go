package gen

import (
	"encoding/json"
	"sync"

	"github.com/google/go-github/v28/github"
)

// See: https://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref
var statusText = `{"url":"https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e","avatar_url":"https://github.com/images/error/hubot_happy.gif","id":1,"node_id":"MDY6U3RhdHVzMQ==","state":"success","description":"Build has completed successfully","target_url":"https://ci.example.com/1000/output","context":"continuous-integration/jenkins","created_at":"2012-07-20T01:19:13Z","updated_at":"2012-07-20T01:19:13Z","creator":{"login":"octocat","id":1,"node_id":"MDQ6VXNlcjE=","avatar_url":"https://github.com/images/error/octocat_happy.gif","gravatar_id":"","url":"https://api.github.com/users/octocat","html_url":"https://github.com/octocat","followers_url":"https://api.github.com/users/octocat/followers","following_url":"https://api.github.com/users/octocat/following{/other_user}","gists_url":"https://api.github.com/users/octocat/gists{/gist_id}","starred_url":"https://api.github.com/users/octocat/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/octocat/subscriptions","organizations_url":"https://api.github.com/users/octocat/orgs","repos_url":"https://api.github.com/users/octocat/repos","events_url":"https://api.github.com/users/octocat/events{/privacy}","received_events_url":"https://api.github.com/users/octocat/received_events","type":"User","site_admin":false}}`
var status *github.RepoStatus
var statusOnce sync.Once

func RepoStatus() (*github.RepoStatus, error) {
	var err error
	statusOnce.Do(func() {
		err = json.Unmarshal([]byte(statusText), &status)
	})
	if err != nil {
		return nil, err
	}
	s := *status
	return &s, nil
}

func RepoStatuses(length int) ([]*github.RepoStatus, error) {
	values := make([]*github.RepoStatus, length)
	for i := 0; i < length; i++ {
		v, err := RepoStatus()
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}
