package action

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	ev "gopkg.in/go-playground/webhooks.v5/github"
)

// https://help.github.com/en/actions/automating-your-workflow-with-github-actions/using-environment-variables
var (
	Home       = os.Getenv("HOME")
	Workflow   = os.Getenv("GITHUB_WORKFLOW")
	Action     = os.Getenv("GITHUB_ACTION")
	Actions    = func(s string) bool { return s == "true" }(os.Getenv("GITHUB_ACTIONS"))
	Actor      = os.Getenv("GITHUB_ACTOR")
	Repository = os.Getenv("GITHUB_REPOSITORY")
	EventName  = os.Getenv("GITHUB_EVENT_NAME")
	EventPath  = os.Getenv("GITHUB_EVENT_PATH")
	Workspace  = os.Getenv("GITHUB_WORKSPACE")
	SHA        = os.Getenv("GITHUB_SHA")
	Ref        = os.Getenv("GITHUB_REF")
	HeadRef    = os.Getenv("GITHUB_HEAD_REF")
	BaseRef    = os.Getenv("GITHUB_BASE_REF")
)

func PullNumber() *int {
	if !strings.HasPrefix(Ref, "refs/pull/") {
		return nil
	}
	s := strings.Split(Ref, "/") // refs/pull/:prNumber/merge
	i, _ := strconv.Atoi(s[2])
	return &i
}

func BranchName() *string {
	ref := Ref
	switch ev.Event(EventName) {
	case ev.CreateEvent:
		return &ref
	case ev.DeploymentEvent:
		if ref == "" {
			return nil
		}
		return &ref
	case ev.DeploymentStatusEvent:
		if ref == "" {
			return nil
		}
		return &ref
	case ev.PushEvent:
		return &ref
	case ev.ReleaseEvent:
		branch := fmt.Sprintf("refs/tags/%s", ref)
		return &branch
	}
	return nil
}

func TagName() *string {
	ref := Ref
	switch ev.Event(EventName) {
	case ev.CreateEvent:
		if !strings.HasPrefix(ref, "refs/tags/") {
			return nil
		}
		tag := strings.TrimPrefix(ref, "refs/tags/")
		return &tag
	case ev.DeploymentEvent:
		if !strings.HasPrefix(ref, "refs/tags/") {
			return nil
		}
		tag := strings.TrimPrefix(ref, "refs/tags/")
		return &tag
	case ev.DeploymentStatusEvent:
		if !strings.HasPrefix(ref, "refs/tags/") {
			return nil
		}
		tag := strings.TrimPrefix(ref, "refs/tags/")
		return &tag
	case ev.ReleaseEvent:
		return &ref
	}
	return nil
}
