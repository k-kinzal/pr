package api

import (
	"context"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type Client struct {
	github *github.Client
}

func NewClient(token string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.TODO(), ts)

	return &Client{
		github: github.NewClient(tc),
	}
}
