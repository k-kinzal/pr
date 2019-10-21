package api

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/google/go-github/v28/github"
	jmespath "github.com/jmespath/go-jmespath"
)

const (
	PerPage = 100
)

type PullRequest struct {
	*github.PullRequest
	Owner    *string              `json:"-"`
	Repo     *string              `json:"-"`
	Statuses []*github.RepoStatus `json:"statuses"`
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
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	rateLimit := 10
	requestNum := 0
	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(rateLimit)), rateLimit)
	getAllPage := func(ch chan interface{}, errCh chan error, limit int, callback func(opt *github.ListOptions) (interface{}, *github.Response, error)) {
		requestNum++
		go func() {
			if err := limiter.Wait(childCtx); err != nil {
				errCh <- err
				return
			}
			opt := &github.ListOptions{
				Page:    0,
				PerPage: PerPage,
			}
			data, resp, err := callback(opt)
			if err != nil {
				errCh <- err
				return
			}
			ch <- data
			if resp.NextPage == -1 {
				return
			}
			allPage := resp.LastPage
			if limit > 0 && allPage > limit/opt.PerPage {
				allPage = limit / opt.PerPage
			}
			for i := 1; i < allPage; i++ {
				requestNum++
				go func(page int) {
					if err := limiter.Wait(childCtx); err != nil {
						errCh <- err
						return
					}
					opt.Page = page
					data, _, err := callback(opt)
					if err != nil {
						errCh <- err
						return
					}
					ch <- data
				}(i)
			}
		}()
	}

	ch := make(chan interface{})
	errCh := make(chan error)

	if num := filter.GetNumber(); num > 0 {
		requestNum++
		go func() {
			pull, _, err := c.github.PullRequests.Get(childCtx, owner, repo, num)
			if err != nil {
				errCh <- err
				return
			}
			ch <- []*github.PullRequest{pull}
		}()
	} else {
		getAllPage(ch, errCh, filter.GetLimit(), func(opt *github.ListOptions) (interface{}, *github.Response, error) {
			pullOptions := &github.PullRequestListOptions{
				State:       filter.GetState(),
				Head:        filter.GetHead(),
				Base:        filter.GetBase(),
				Sort:        "created",
				Direction:   "desc",
				ListOptions: *opt,
			}
			return c.github.PullRequests.List(childCtx, owner, repo, pullOptions)
		})
	}

	pulls := make([]*github.PullRequest, 0)
	statusesMap := make(map[string][]*github.RepoStatus, 0)
	for i := 0; i < requestNum; i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			break
		case data := <-ch:
			switch reflect.TypeOf(data).Kind() {
			case reflect.Slice, reflect.Array:
				v := reflect.ValueOf(data)
				for i := 0; i < v.Len(); i++ {
					switch vv := v.Index(i).Interface().(type) {
					case *github.PullRequest:
						pulls = append(pulls, vv)
						func(pull *github.PullRequest) {
							getAllPage(ch, errCh, -1, func(opt *github.ListOptions) (interface{}, *github.Response, error) {
								return c.github.Repositories.ListStatuses(childCtx, owner, repo, pull.GetHead().GetSHA(), opt)
							})
						}(vv)
					case *github.RepoStatus:
						s := strings.Split(vv.GetURL(), "/")
						statusesMap[s[len(s)-1]] = append(statusesMap[s[len(s)-1]], vv)
					}
				}
			}
		case err, _ := <-errCh:
			cancel()
			return nil, err
		}
	}

	var allPulls []*PullRequest
	for _, pull := range pulls {
		statuses, ok := statusesMap[pull.GetHead().GetRef()]
		if !ok {
			statuses = make([]*github.RepoStatus, 0)
		}
		allPulls = append(allPulls, &PullRequest{
			PullRequest: pull,
			Owner:       &owner,
			Repo:        &repo,
			Statuses:    statuses,
		})
	}

	return filter.Apply(allPulls)
}
