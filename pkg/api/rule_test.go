package api_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/k-kinzal/pr/pkg/api"
)

func TestPullRequestRules_GetNumberEQ(t *testing.T) {
	r := []string{
		"number == `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberNQ(t *testing.T) {
	r := []string{
		"number != `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberNE(t *testing.T) {
	r := []string{
		"number != `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberGT(t *testing.T) {
	r := []string{
		"number > `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberGE(t *testing.T) {
	r := []string{
		"number >= `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberLT(t *testing.T) {
	r := []string{
		"number < `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetNumberLE(t *testing.T) {
	r := []string{
		"number <= `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetNumber() != 1 {
		t.Errorf("expect `1`, but actual `%#v`", rules.GetNumber())
	}
}

func TestPullRequestRules_GetStateEQ(t *testing.T) {
	r := []string{
		"state == `\"open\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetState() != "open" {
		t.Errorf("expect `open`, but actual `%#v`", rules.GetState())
	}
}

func TestPullRequestRules_GetStateNE(t *testing.T) {
	r := []string{
		"state != `\"open\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetState() != "open" {
		t.Errorf("expect `open`, but actual `%#v`", rules.GetState())
	}
}

func TestPullRequestRules_GetStateDefault(t *testing.T) {
	r := []string{}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetState() != "all" {
		t.Errorf("expect `all`, but actual `%#v`", rules.GetState())
	}
}

func TestPullRequestRules_GetHeadEQ(t *testing.T) {
	r := []string{
		"head == `\"branch-name\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetHead() != "branch-name" {
		t.Errorf("expect `branch-name`, but actual `%#v`", rules.GetHead())
	}
}

func TestPullRequestRules_GetHeadNE(t *testing.T) {
	r := []string{
		"head != `\"branch-name\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetHead() != "branch-name" {
		t.Errorf("expect `branch-name`, but actual `%#v`", rules.GetHead())
	}
}

func TestPullRequestRules_GetBaseEQ(t *testing.T) {
	r := []string{
		"base == `\"branch-name\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetBase() != "branch-name" {
		t.Errorf("expect `branch-name`, but actual `%#v`", rules.GetBase())
	}
}

func TestPullRequestRules_GetBaseNE(t *testing.T) {
	r := []string{
		"base != `\"branch-name\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetBase() != "branch-name" {
		t.Errorf("expect `branch-name`, but actual `%#v`", rules.GetBase())
	}
}

func TestPullRequestRules_GetLimit(t *testing.T) {
	r := []string{}
	rules := api.NewPullRequestRules(r, 100)
	if rules.GetLimit() != 100 {
		t.Errorf("expect `100`, but actual `%#v`", rules.GetLimit())
	}
}

func TestPullRequestRules_GetRule(t *testing.T) {
	r := []string{
		"state == `\"open\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	if len(rules.GetRules()) != 1 || rules.GetRules()[0] != "state == `\"open\"`" {
		t.Errorf("expect `state == `\"open\"``, but actual `%#v`", rules.GetRules())
	}
}

func TestPullRequestRules_Expression(t *testing.T) {
	api.Now = time.Now().UTC().Unix()
	r := []string{
		"state == `\"open\"`",
		"now() > `\"2006-01-02T15:04:05Z\"`",
		"now() > `\"15:04:05\"`",
	}
	rules := api.NewPullRequestRules(r, 100)
	timeString := fmt.Sprintf("%sT15:04:05Z", time.Unix(api.Now, 0).UTC().Format("2006-01-02"))
	tm, _ := time.Parse("2006-01-02T15:04:05Z", timeString)
	expect := fmt.Sprintf("[?state == `\"open\"`] | [?`%d` > `1136214245`] | [?`%d` > `%d`]", api.Now, api.Now, tm.UTC().Unix())
	if rules.Expression() != expect {
		t.Errorf("expect `%s`, but actual `%#v`", expect, rules.Expression())
	}
}

func ptri64(i int64) *int64 {
	return &i
}

func ptrs(s string) *string {
	return &s
}

func TestPullRequestRules_Apply(t *testing.T) {
	now := &api.Timestamp{Time: time.Now()}
	pulls := []*api.PullRequest{
		{
			ID:        ptri64(1),
			State:     ptrs("open"),
			CreatedAt: now,
			Statuses: []*api.RepoStatus{
				{
					State: ptrs("success"),
				},
				{
					State: ptrs("pending"),
				},
			},
			Owner: ptrs("example"),
			Repo:  ptrs("repo"),
		},
	}
	r := []string{
		"state == `\"open\"`",
		"length(statuses) == `2`",
		"length(statuses[?state == `\"success\"`]) == `1`",
		"length(statuses[?state == `\"pending\"`]) == `1`",
	}
	rules := api.NewPullRequestRules(r, 100)
	filtered, err := rules.Apply(pulls)
	if err != nil {
		t.Error(err)
		return
	}
	if len(filtered) != 1 {
		t.Errorf("expect `1`, but actual `%d`: %v", len(filtered), filtered)
	}
}
