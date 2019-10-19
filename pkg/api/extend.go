package api

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"strings"
	"time"
)

type Status struct {
	Url         *string      `json:"url,omitempty"`
	AvatarUrl   *string      `json:"avatar_url,omitempty"`
	Id          *int64       `json:"id,omitempty"`
	NodeId      *string      `json:"node_id,omitempty"`
	State       *string      `json:"state,omitempty"`
	Description *string      `json:"description,omitempty"`
	TargetUrl   *string      `json:"target_url,omitempty"`
	Context     *string      `json:"context,omitempty"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
	Creator     *github.User `json:"creator,omitempty"`
}

func (s Status) String() string {
	return github.Stringify(s)
}

type StatusService struct {
	client *github.Client
}

type StatusCreateOptions struct {
	State       *string `json:"state,omitempty"`
	TargetURL   *string `json:"target_url,omitempty"`
	Description *string `json:"description,omitempty"`
	Context     *string `json:"context,omitempty"`
}

func (s *StatusService) Create(ctx context.Context, owner string, repo string, sha string, opts *StatusCreateOptions) (*Status, *github.Response, error) {
	return nil, nil, fmt.Errorf("StatusService::Create is not implemented")
}

func (s *StatusService) List(ctx context.Context, owner string, repo string, ref string) ([]*Status, *github.Response, error) {
	u := fmt.Sprintf("/repos/%s/%s/commits/%s/statuses", owner, repo, ref)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		panic(err)
	}

	acceptHeaders := []string{
		"application/vnd.github.symmetra-preview+json",
		"application/vnd.github.sailor-v-preview+json",
		"application/vnd.github.shadow-cat-preview+json",
	}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))

	var statuses []*Status
	res, err := s.client.Do(ctx, req, &statuses)
	if err != nil {
		return nil, res, err
	}

	return statuses, res, nil
}

func (s *StatusService) Get(ctx context.Context, owner string, repo string, ref string) ([]*Status, *github.Response, error) {
	return nil, nil, fmt.Errorf("StatusService::Get is not implemented")
}
