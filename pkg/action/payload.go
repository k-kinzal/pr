package action

import (
	"encoding/json"
	"io/ioutil"

	ev "gopkg.in/go-playground/webhooks.v5/github"
)

// https://help.github.com/ja/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows
var Payload interface{}

func init() {
	bytes, err := ioutil.ReadFile(EventPath)
	if err != nil {
		return
	}
	switch ev.Event(EventName) {
	case ev.CheckRunEvent:
		var v *ev.CheckRunPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.CheckSuiteEvent:
		var v *ev.CheckRunPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.CreateEvent:
		var v *ev.CreatePayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.DeleteEvent:
		var v *ev.DeletePayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.DeploymentEvent:
		var v *ev.DeploymentPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.DeploymentStatusEvent:
		var v *ev.DeploymentStatusPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.ForkEvent:
		var v *ev.ForkPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.GollumEvent:
		var v *ev.GollumPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.IssueCommentEvent:
		var v *ev.IssueCommentPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.IssuesEvent:
		var v *ev.IssuesPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.LabelEvent:
		var v *ev.LabelPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.MemberEvent:
		var v *ev.MemberPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.MilestoneEvent:
		var v *ev.MilestonePayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.PageBuildEvent:
		var v *ev.PageBuildPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.ProjectEvent:
		var v *ev.CheckRunPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.ProjectCardEvent:
		var v *ev.ProjectCardPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.ProjectColumnEvent:
		var v *ev.ProjectColumnPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
		return
	case ev.PublicEvent:
		var v *ev.PublicPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.PullRequestEvent:
		var v *ev.PullRequestPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.PullRequestReviewEvent:
		var v *ev.PullRequestReviewPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.PullRequestReviewCommentEvent:
		var v *ev.PullRequestReviewCommentPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.PushEvent:
		var v *ev.PushPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.ReleaseEvent:
		var v *ev.ReleasePayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.StatusEvent:
		var v *ev.StatusPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	case ev.WatchEvent:
		var v *ev.WatchPayload
		if err := json.Unmarshal(bytes, &v); err == nil {
			Payload = v
		}
	}
}
