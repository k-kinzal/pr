package api

import (
	"context"
	"time"

	"github.com/k-kinzal/pr/pkg/httpratelimit"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

type Options struct {
	Token     string
	RateLimit int
}

type Client struct {
	github *github.Client
}

func NewClient(ctx context.Context, opt *Options) *Client {
	// Create HTTP RateLimit client
	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(opt.RateLimit)), opt.RateLimit)
	client := httpratelimit.NewClient(ctx, limiter)

	// Create OAuth2 client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: opt.Token},
	)
	tc := oauth2.NewClient(context.WithValue(ctx, oauth2.HTTPClient, client), ts)

	// Create Github API client
	return &Client{
		github: github.NewClient(tc),
	}
}
