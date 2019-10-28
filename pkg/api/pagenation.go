package api

import (
	"context"
	"math"
	"reflect"
	"sync"

	"github.com/google/go-github/v28/github"
)

const (
	PerPage = 100
)

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
	ch chan *PagenationResult
}

func (p *Pagenation) Request(ctx context.Context, callback func(opt *github.ListOptions) (interface{}, *github.Response, error)) {
	p.RequestWithLimit(ctx, math.MaxInt64, callback)
}

func (p *Pagenation) RequestWithLimit(ctx context.Context, maxPage int, callback func(opt *github.ListOptions) (interface{}, *github.Response, error)) {
	wg := new(sync.WaitGroup)

	receiver := p.ch
	sender := make(chan *PagenationResult)
	p.ch = sender
	if receiver != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range receiver {
				sender <- v
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		r := func(o *github.ListOptions) (interface{}, *github.Response, error) {
			data, resp, err := callback(o)
			switch reflect.TypeOf(data).Kind() {
			case reflect.Slice, reflect.Array:
				v := reflect.ValueOf(data)
				for i := 0; i < v.Len(); i++ {
					sender <- &PagenationResult{v.Index(i).Interface(), resp, err}
				}
			default:
				sender <- &PagenationResult{data, resp, err}
			}
			return data, resp, err
		}
		o := &github.ListOptions{
			Page:    0,
			PerPage: PerPage,
		}
		_, resp, err := r(o)
		if err != nil {
			return
		}

		lastPage := resp.LastPage
		if lastPage > maxPage {
			lastPage = maxPage
		}
		for i := 1; i < lastPage; i++ {
			wg.Add(1)
			go func(page int) {
				defer wg.Done()
				o.Page = page
				r(o)
			}(i)
		}
	}()

	go func() {
		wg.Wait()
		close(sender)
	}()
}

func (p *Pagenation) Done() <-chan *PagenationResult {
	return p.ch
}
