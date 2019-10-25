package httpratelimit

import (
	"context"
	"net/http"

	"golang.org/x/time/rate"
)

type Transport struct {
	Limiter *rate.Limiter
	Base    http.RoundTripper
}

func (t *Transport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	if err := t.Limiter.Wait(ctx); err != nil {
		return nil, err
	}
	return t.base().RoundTrip(req)
}

var HTTPClient struct{}

func contextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
			return hc
		}
	}
	return http.DefaultClient
}

func NewClient(ctx context.Context, l *rate.Limiter) *http.Client {
	if l == nil {
		return contextClient(ctx)
	}
	return &http.Client{
		Transport: &Transport{
			Base:    contextClient(ctx).Transport,
			Limiter: l,
		},
	}
}
