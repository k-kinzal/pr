package action

import (
	"encoding/json"
	"io/ioutil"

	"github.com/google/go-github/v28/github"
)

// https://help.github.com/ja/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows
var Payload interface{}

func init() {
	bytes, err := ioutil.ReadFile(EventPath)
	if err != nil {
		return
	}
	switch EventName {
	case "check_run":
		var v *github.CheckRunEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "check_suite":
		var v *github.CheckRunEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "create":
		var v *github.CreateEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "delete":
		var v *github.DeleteEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "deployment":
		var v *github.DeploymentEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "deployment_status":
		var v *github.DeploymentStatusEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "fork":
		var v *github.ForkEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "gollum":
		var v *github.GollumEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "issue_comment":
		var v *github.IssueCommentEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "issues":
		var v *github.IssuesEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "label":
		var v *github.LabelEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "member":
		var v *github.MemberEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "milestone":
		var v *github.MilestoneEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "page_build":
		var v *github.PageBuildEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "project":
		var v *github.CheckRunEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "project_card":
		var v *github.ProjectCardEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "project_column":
		var v *github.ProjectColumnEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
		return
	case "public":
		var v *github.PublicEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "pull_request":
		var v *github.PullRequestEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "pull_request_review":
		var v *github.PullRequestReviewEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "pull_request_review_comment":
		var v *github.PullRequestReviewCommentEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "push":
		var v *github.PushEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "release":
		var v *github.ReleaseEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "status":
		var v *github.StatusEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case "watch":
		var v *github.WatchEvent
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	}
}
