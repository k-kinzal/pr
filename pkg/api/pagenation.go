package api

import (
	"context"
	"math"
	"sync"

	"github.com/google/go-github/v28/github"
)

const (
	PerPage = 100
)

type Counter struct {
	n int
	m sync.Mutex
}

func (c *Counter) Increment() {
	c.m.Lock()
	c.n++
	c.m.Unlock()
}

func (c *Counter) Num() int {
	return c.n
}

type PagenationResult struct {
	i interface{}
	r *github.Response
	e error
}

func (pr *PagenationResult) Interface() interface{} {
	return pr.i
}

func (pr *PagenationResult) Response() *github.Response {
	return pr.r
}

func (pr *PagenationResult) Error() error {
	return pr.e
}

type Pagenation struct {
	ch  chan *PagenationResult
	wg  *sync.WaitGroup
	cnt Counter
}

func (p *Pagenation) Request(ctx context.Context, callback func(opt *github.ListOptions) (interface{}, *github.Response, error)) {
	p.RequestWithLimit(ctx, math.MaxInt64, callback)
}

func (p *Pagenation) RequestWithLimit(ctx context.Context, maxPage int, callback func(opt *github.ListOptions) (interface{}, *github.Response, error)) {
	p.cnt.Increment()
	go func() {
		o := &github.ListOptions{
			Page:    0,
			PerPage: PerPage,
		}
		data, resp, err := callback(o)
		p.ch <- &PagenationResult{data, resp, err}
		if err != nil {
			return
		}

		lastPage := resp.LastPage
		if lastPage > maxPage {
			lastPage = maxPage
		}
		for i := 1; i < lastPage; i++ {
			p.cnt.Increment()
			go func(page int) {
				o.Page = page
				data, resp, err := callback(o)
				p.ch <- &PagenationResult{data, resp, err}
			}(i)
		}
	}()

}

func (p *Pagenation) Done() <-chan *PagenationResult {
	return p.ch
}

func (p *Pagenation) RequestedNum() int {
	return p.cnt.Num()
}

func NewPagenation() *Pagenation {
	return &Pagenation{
		ch: make(chan *PagenationResult),
		wg: new(sync.WaitGroup),
		cnt: Counter{
			n: 0,
			m: sync.Mutex{},
		},
	}
}
