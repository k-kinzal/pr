package api

import (
	"context"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type Client struct {
	github struct {
		*github.Client
		Statuses StatusService
	}
}

func NewClient(token string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.TODO(), ts)
	cli := github.NewClient(tc)

	return &Client{
		github: struct {
			*github.Client
			Statuses StatusService
		}{
			Client:   cli,
			Statuses: StatusService{cli},
		},
	}
}
