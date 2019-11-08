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
	dateRegexp        = regexp.MustCompile(`"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z"`)
	timeRegexp        = regexp.MustCompile(`"\d{2}:\d{2}:\d{2}"`)
	nowFunctionRegexp = regexp.MustCompile(`now\(\)`)
)

type PullRequestRules struct {
	rules []string
	limit int

	number int
	state  string
	head   string
	base   string
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
			if data[i].ID == vv.ID {
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
	now := fmt.Sprintf("`%d`", Now)

	var number int
	var state string
	var head string
	var base string
	for i, rule := range rules {
		{
			m := numberRegexp.FindSubmatch([]byte(rule))
			if len(m) == 2 {
				n, _ := strconv.Atoi(string(m[1]))
				number = n
			}
		}
		{
			m := stateRegexp.FindSubmatch([]byte(rule))
			if len(m) == 2 {
				state = string(m[1])
			}
		}
		{
			m := headRegexp.FindSubmatch([]byte(rule))
			if len(m) == 2 {
				head = string(m[1])
			}
		}
		{
			m := baseRegexp.FindSubmatch([]byte(rule))
			if len(m) == 2 {
				base = string(m[1])
			}
		}

		rule = nowFunctionRegexp.ReplaceAllString(rule, now)

		rule = string(dateRegexp.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
			t, err := time.Parse("2006-01-02T15:04:05Z", strings.Trim(string(bytes), "\""))
			if err != nil {
				return bytes
			}
			return []byte(fmt.Sprintf("%d", t.UTC().Unix()))
		}))

		rule = string(timeRegexp.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
			s := fmt.Sprintf("%sT%sZ", time.Now().UTC().Format("2006-01-02"), strings.Trim(string(bytes), "\""))
			t, err := time.Parse(dateLayout, s)
			if err != nil {
				return bytes
			}
			return []byte(fmt.Sprintf("%d", t.UTC().Unix()))
		}))

		rules[i] = rule
	}

	return &PullRequestRules{
		rules:  rules,
		limit:  limit,
		number: number,
		state:  state,
		head:   head,
		base:   base,
	}
}
