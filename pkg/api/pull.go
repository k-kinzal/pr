package api

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v28/github"
	jmespath "github.com/jmespath/go-jmespath"
)

const (
	PerPage = 100
)

type PullRequest struct {
	*github.PullRequest
	Owner    *string   `json:"-"`
	Repo     *string   `json:"-"`
	Statuses []*Status `json:"statuses,omitempty"`
}

func (pr *PullRequest) GetOwner() string {
	if pr.Owner == nil {
		return ""
	} else {
		return *pr.Owner
	}
}

func (pr *PullRequest) GetRepo() string {
	if pr.Repo == nil {
		return ""
	} else {
		return *pr.Repo
	}
}

type PullsFilter struct {
	rules []string
	limit int
}

func (f *PullsFilter) GetNumber() int {
	r, _ := regexp.Compile("number ?== ?`([0-9]*?)`")
	for _, rule := range f.rules {
		m := r.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			n, _ := strconv.Atoi(string(m[1]))
			return n
		}
	}
	return -1
}

func (f *PullsFilter) GetState() string {
	r, _ := regexp.Compile("state ?== ?`\"(.*?)\"`")
	for _, rule := range f.rules {
		m := r.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			return string(m[1])
		}
	}
	return "all"
}

func (f *PullsFilter) GetHead() string {
	r, _ := regexp.Compile("head ?== ?`\"(.*?)\"`")
	for _, rule := range f.rules {
		m := r.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			return string(m[1])
		}
	}
	return ""
}

func (f *PullsFilter) GetBase() string {
	r, _ := regexp.Compile("base ?== ?`\"(.*?)\"`")
	for _, rule := range f.rules {
		m := r.FindSubmatch([]byte(rule))
		if len(m) == 2 {
			return string(m[1])
		}
	}
	return ""
}

func (f *PullsFilter) GetLimit() int {
	return f.limit
}

func (f *PullsFilter) GetRules() []string {
	return f.rules
}

func (f *PullsFilter) Expression() string {
	var expressions []string
	for _, rule := range f.rules {
		expressions = append(expressions, fmt.Sprintf("[?%s]", rule))
	}

	exp := "[*]"
	if len(expressions) > 0 {
		exp = strings.Join(expressions, " | ")
	}

	return exp
}

func (f *PullsFilter) Apply(data []*PullRequest) ([]*PullRequest, error) {
	out, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r := regexp.MustCompile(`"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z"`)
	replaced := r.ReplaceAllFunc(out, func(bytes []byte) []byte {
		t, err := time.Parse("2006-01-02T15:04:05Z", strings.Trim(string(bytes), "\""))
		if err != nil {
			return bytes
		}
		return []byte(fmt.Sprintf("%d", t.UTC().Unix()))
	})

	var v interface{}
	if err := json.Unmarshal(replaced, &v); err != nil {
		return nil, err
	}
	fmt.Println(f.Expression())
	filtered, err := jmespath.Search(f.Expression(), v)
	if err != nil {
		return nil, fmt.Errorf("jmespath: %s", err)
	}

	i := 0
	var pulls []*PullRequest
	for _, v := range filtered.([]interface{}) {
		vv, _ := v.(map[string]interface{})
		for ; i < len(data); i++ {
			if data[i].GetID() == int64(vv["id"].(float64)) {
				pulls = append(pulls, data[i])
				break
			}
		}
	}

	if len(pulls) > f.limit {
		return pulls[:f.limit], nil
	} else {
		return pulls, nil
	}

}

func NewPullsFilter(rules []string, limit int) *PullsFilter {
	now := fmt.Sprintf("`%d`", time.Now().UTC().Unix())
	r1 := regexp.MustCompile(`now\(\)`)
	r2 := regexp.MustCompile(`"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z"`)
	r3 := regexp.MustCompile(`"\d{2}:\d{2}:\d{2}"`)
	for i, rule := range rules {
		rule = r1.ReplaceAllString(rule, now)
		rule = string(r2.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
			t, err := time.Parse("2006-01-02T15:04:05Z", strings.Trim(string(bytes), "\""))
			if err != nil {
				return bytes
			}
			return []byte(fmt.Sprintf("%d", t.UTC().Unix()))
		}))
		rule = string(r3.ReplaceAllFunc([]byte(rule), func(bytes []byte) []byte {
			s := fmt.Sprintf("%sT%sZ", time.Now().UTC().Format("2006-01-02"), strings.Trim(string(bytes), "\""))
			t, err := time.Parse("2006-01-02T15:04:05Z", s)
			if err != nil {
				return bytes
			}
			return []byte(fmt.Sprintf("%d", t.UTC().Unix()))
		}))
		rules[i] = rule
	}

	return &PullsFilter{
		rules: rules,
		limit: limit,
	}
}

func (c *Client) GetPulls(ctx context.Context, owner string, repo string, filter *PullsFilter) ([]*PullRequest, error) {
	conv := func(ctx context.Context, pulls []*github.PullRequest) ([]*PullRequest, error) {
		childCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		ch := make(chan *PullRequest)
		errCh := make(chan error)
		for _, pull := range pulls {
			go func() {
				statuses, _, err := c.github.Statuses.List(childCtx, owner, repo, pull.GetHead().GetRef())
				if err != nil {
					errCh <- err
					return
				}
				ch <- &PullRequest{
					PullRequest: pull,
					Statuses:    statuses,
				}
			}()
		}
		var allPulls []*PullRequest
		for i := 0; i < len(pulls); i++ {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					return nil, err
				}
				break
			case pull, _ := <-ch:
				allPulls = append(allPulls, pull)
			case err, _ := <-errCh:
				cancel()
				return nil, err
			}
		}
		return allPulls, nil
	}

	opt := &github.PullRequestListOptions{
		State:       filter.GetState(),
		Head:        filter.GetHead(),
		Base:        filter.GetBase(),
		Sort:        "created",
		Direction:   "desc",
		ListOptions: github.ListOptions{PerPage: PerPage},
	}
	nextPage := -1
	allPage := 0

	var pulls []*github.PullRequest
	if num := filter.GetNumber(); num > 0 {
		pull, _, err := c.github.PullRequests.Get(ctx, owner, repo, num)
		if err != nil {
			return nil, err
		}
		pulls = append(pulls, pull)
		nextPage = -1
		allPage = 1
	} else {
		p, resp, err := c.github.PullRequests.List(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}
		pulls = append(pulls, p...)
		nextPage = resp.NextPage
		allPage = resp.LastPage
	}

	if nextPage == -1 || filter.GetLimit() > PerPage {
		p, err := conv(ctx, pulls)
		if err != nil {
			return nil, err
		}
		return filter.Apply(p)
	}

	if allPage > (filter.GetLimit() / PerPage) {
		allPage = filter.GetLimit() / PerPage
	}

	allPulls, err := conv(ctx, pulls)
	if err != nil {
		return nil, err
	}

	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan []*PullRequest)
	errCh := make(chan error)
	for page := 1; page < allPage; page++ {
		go func() {
			opt.Page = page
			pulls, _, err := c.github.PullRequests.List(childCtx, owner, repo, opt)
			if err != nil {
				errCh <- err
				return
			}
			p, err := conv(ctx, pulls)
			if err != nil {
				errCh <- err
				return
			}
			ch <- p
		}()
	}

	for i := 1; i < allPage; i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			break
		case pulls, _ := <-ch:
			allPulls = append(allPulls, pulls...)
		case err, _ := <-errCh:
			cancel()
			return nil, err
		}
	}

	return filter.Apply(allPulls)
}
