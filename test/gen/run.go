package gen

import (
	"encoding/json"
	"sync"

	"github.com/google/go-github/v28/github"
)

// See: https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-specific-ref
var runText = `{"total_count":1,"check_runs":[{"id":4,"head_sha":"ce587453ced02b1526dfb4cb910479d431683101","node_id":"MDg6Q2hlY2tSdW40","external_id":"","url":"https://api.github.com/repos/github/hello-world/check-runs/4","html_url":"http://github.com/github/hello-world/runs/4","details_url":"https://example.com","status":"completed","conclusion":"neutral","started_at":"2018-05-04T01:14:52Z","completed_at":"2018-05-04T01:14:52Z","output":{"title":"Mighty Readme report","summary":"There are 0 failures, 2 warnings, and 1 notice.","text":"You may have some misspelled words on lines 2 and 4. You also may want to add a section in your README about how to install your app.","annotations_count":2,"annotations_url":"https://api.github.com/repos/github/hello-world/check-runs/4/annotations"},"name":"mighty_readme","check_suite":{"id":5},"app":{"id":1,"slug":"octoapp","node_id":"MDExOkludGVncmF0aW9uMQ==","owner":{"login":"github","id":1,"node_id":"MDEyOk9yZ2FuaXphdGlvbjE=","url":"https://api.github.com/orgs/github","repos_url":"https://api.github.com/orgs/github/repos","events_url":"https://api.github.com/orgs/github/events","hooks_url":"https://api.github.com/orgs/github/hooks","issues_url":"https://api.github.com/orgs/github/issues","members_url":"https://api.github.com/orgs/github/members{/member}","public_members_url":"https://api.github.com/orgs/github/public_members{/member}","avatar_url":"https://github.com/images/error/octocat_happy.gif","description":"A great organization"},"name":"Octocat App","description":"","external_url":"https://example.com","html_url":"https://github.com/apps/octoapp","created_at":"2017-07-08T16:18:44-04:00","updated_at":"2017-07-08T16:18:44-04:00","permissions":{"metadata":"read","contents":"read","issues":"write","single_file":"write"},"events":["push","pull_request"]},"pull_requests":[{"url":"https://api.github.com/repos/github/hello-world/pulls/1","id":1934,"number":3956,"head":{"ref":"say-hello","sha":"3dca65fa3e8d4b3da3f3d056c59aee1c50f41390","repo":{"id":526,"url":"https://api.github.com/repos/github/hello-world","name":"hello-world"}},"base":{"ref":"master","sha":"e7fdf7640066d71ad16a86fbcbb9c6a10a18af4f","repo":{"id":526,"url":"https://api.github.com/repos/github/hello-world","name":"hello-world"}}}]}]}`
var run *github.ListCheckRunsResults
var runOnce sync.Once

func ListCheckRunsResults() (*github.ListCheckRunsResults, error) {
	var err error
	runOnce.Do(func() {
		err = json.Unmarshal([]byte(runText), &run)
	})
	if err != nil {
		return nil, err
	}
	r := *run
	return &r, nil
}
