package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/xerrors"

	jmespath "github.com/jmespath/go-jmespath"
)

var (
	Now               = time.Now().UTC().Unix()
	dateLayout        = "2006-01-02T15:04:05Z"
	numberRegexp      = regexp.MustCompile("number\\s*[=!<>]=?\\s*`([0-9]*?)`")
	stateRegexp       = regexp.MustCompile("state\\s*[=!]=\\s*`\"(.*?)\"`")
	headRegexp        = regexp.MustCompile("head\\s*[=!]=\\s*`\"(.*?)\"`")
	baseRegexp        = regexp.MustCompile("base\\s*[=!]=\\s*`\"(.*?)\"`")
	shaRegexp         = regexp.MustCompile("head.sha\\s*[=!]=\\s*`\"(.*?)\"`")
	dateRegexp        = regexp.MustCompile("`\"\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z\"`")
	timeRegexp        = regexp.MustCompile("`\"\\d{2}:\\d{2}:\\d{2}\"`")
	nowFunctionRegexp = regexp.MustCompile(`now\(\)`)
)

type PullRequestRules struct {
	rules []string
	limit int

	number int
	state  string
	head   string
	base   string

	sha string
}

func (r *PullRequestRules) Add(rule string) {
	now := fmt.Sprintf("`%d`", Now)
	{
		m := numberRegexp.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			n, _ := strconv.Atoi(string(m[1]))
			r.number = n
		}
	}
	{
		m := stateRegexp.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			r.state = string(m[1])
		}
	}
	{
		m := headRegexp.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			r.head = string(m[1])
		}
	}
	{
		m := baseRegexp.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			r.base = string(m[1])
		}
	}
	{
		m := shaRegexp.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			r.sha = string(m[1])
		}
	}

	rule = nowFunctionRegexp.ReplaceAllString(rule, now)

	rule = string(dateRegexp.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
		t, err := time.Parse("2006-01-02T15:04:05Z", strings.Trim(string(bytes), "`\""))
		if err != nil {
			return bytes
		}
		return []byte(fmt.Sprintf("`%d`", t.UTC().Unix()))
	}))

	rule = string(timeRegexp.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
		s := fmt.Sprintf("%sT%sZ", time.Now().UTC().Format("2006-01-02"), strings.Trim(string(bytes), "`\""))
		t, err := time.Parse(dateLayout, s)
		if err != nil {
			return bytes
		}
		return []byte(fmt.Sprintf("`%d`", t.UTC().Unix()))
	}))

	r.rules = append(r.rules, rule)
}

func (r *PullRequestRules) GetNumber() int {
	return r.number
}

func (r *PullRequestRules) GetState() string {
	if r.state == "" {
		return "all"
	}
	return r.state
}

func (r *PullRequestRules) GetHead() string {
	return r.head
}

func (r *PullRequestRules) GetBase() string {
	return r.base
}

func (r *PullRequestRules) GetSHA() string {
	return r.sha
}

func (r *PullRequestRules) GetLimit() int {
	if r.limit <= 0 {
		return 100
	}
	return r.limit
}

func (r *PullRequestRules) GetRules() []string {
	return r.rules
}

func (r *PullRequestRules) Expression() string {
	var expressions []string
	for _, rule := range r.GetRules() {
		expressions = append(expressions, fmt.Sprintf("[?%s]", rule))
	}

	exp := "[*]"
	if len(expressions) > 0 {
		exp = strings.Join(expressions, " | ")
	}

	return exp
}

func (r *PullRequestRules) SearchRules() *PullRequestRules {
	rules := make([]string, 0)
	if number := r.GetNumber(); number > 0 {
		rules = append(rules, fmt.Sprintf("number == `%d`", number))
	}
	if branch := r.GetHead(); branch != "" {
		rules = append(rules, fmt.Sprintf("head.ref == `\"%s\"`", branch))
	}
	if sha := r.GetSHA(); sha != "" {
		rules = append(rules, fmt.Sprintf("head.sha == `\"%s\"`", sha))
	}
	if state := r.GetState(); state != "" {
		rules = append(rules, fmt.Sprintf("state == `\"%s\"`", state))
	}
	return NewPullRequestRules(rules, r.GetLimit())
}

func (r *PullRequestRules) Apply(data []*PullRequest) ([]*PullRequest, error) {
	filtered, err := jmespath.Search(r.Expression(), data)
	if err != nil {
		return nil, xerrors.Errorf("jmespath: %s", err)
	}
	if filtered == nil {
		return make([]*PullRequest, 0), nil
	}

	i := 0
	var pulls []*PullRequest
	for _, v := range filtered.([]interface{}) {
		vv, _ := v.(*PullRequest)
		for ; i < len(data); i++ {
			if data[i].Id == vv.Id {
				pulls = append(pulls, data[i])
				break
			}
		}
	}

	if limit := r.GetLimit(); len(pulls) > limit {
		return pulls[:limit], nil
	} else {
		return pulls, nil
	}
}

func NewPullRequestRules(rules []string, limit int) *PullRequestRules {
	r := &PullRequestRules{}
	for _, rule := range rules {
		r.Add(rule)
	}
	r.limit = limit
	return r
}
