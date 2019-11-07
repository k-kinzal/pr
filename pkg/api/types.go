package api

import (
	"strconv"
	"time"

	"github.com/google/go-github/v28/github"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.UTC().Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		t.Time = time.Unix(i, 0)
	} else {
		t.Time, err = time.Parse(`"`+time.RFC3339+`"`, str)
	}
	return err
}

func newTimestamp(t *time.Time) *Timestamp {
	if t == nil {
		return nil
	}
	return &Timestamp{Time: *t}
}

type Team struct {
	ID              *int64  `json:"id,omitempty"`
	NodeID          *string `json:"node_id,omitempty"`
	URL             *string `json:"url,omitempty"`
	Name            *string `json:"name,omitempty"`
	Slug            *string `json:"slug,omitempty"`
	Description     *string `json:"description,omitempty"`
	Privacy         *string `json:"privacy,omitempty"`
	Permission      *string `json:"permission,omitempty"`
	MembersURL      *string `json:"members_url,omitempty"`
	RepositoriesURL *string `json:"repositories_url,omitempty"`
	Parent          *Team   `json:"parent,omitempty"`
}

func newTeam(team *github.Team) *Team {
	if team == nil {
		return nil
	}
	t := &Team{}
	t.ID = team.ID
	t.NodeID = team.NodeID
	t.URL = team.URL
	//t.HTMLURL = team.HTMLURL // go-github not support
	t.Name = team.Name
	t.Slug = team.Slug
	t.Description = team.Description
	t.Privacy = team.Privacy
	t.Permission = team.Permission
	t.MembersURL = team.MembersURL
	t.RepositoriesURL = team.RepositoriesURL
	t.Parent = newTeam(team.Parent)
	return t
}

type User struct {
	Login             *string `json:"login,omitempty"`
	ID                *int64  `json:"id,omitempty"`
	NodeID            *string `json:"node_id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin,omitempty"`
}

func newUser(user *github.User) *User {
	if user == nil {
		return nil
	}
	u := &User{}
	u.Login = user.Login
	u.ID = user.ID
	u.NodeID = user.NodeID
	u.AvatarURL = user.AvatarURL
	u.GravatarID = user.GravatarID
	u.URL = user.URL
	u.HTMLURL = user.HTMLURL
	u.FollowersURL = user.FollowersURL
	u.FollowingURL = user.FollowingURL
	u.GistsURL = user.GistsURL
	u.StarredURL = user.StarredURL
	u.SubscriptionsURL = user.SubscriptionsURL
	u.OrganizationsURL = user.OrganizationsURL
	u.ReposURL = user.ReposURL
	u.EventsURL = user.EventsURL
	u.ReceivedEventsURL = user.ReceivedEventsURL
	u.Type = user.Type
	u.SiteAdmin = user.SiteAdmin
	return u
}

type Repository struct {
	ID                 *int64           `json:"id,omitempty"`
	NodeID             *string          `json:"node_id,omitempty"`
	Name               *string          `json:"name,omitempty"`
	FullName           *string          `json:"full_name,omitempty"`
	Owner              *User            `json:"owner,omitempty"`
	Private            *bool            `json:"private,omitempty"`
	HTMLURL            *string          `json:"html_url,omitempty"`
	Description        *string          `json:"description,omitempty"`
	Fork               *bool            `json:"fork,omitempty"`
	URL                *string          `json:"url,omitempty"`
	ArchiveURL         *string          `json:"archive_url,omitempty"`
	AssigneesURL       *string          `json:"assignees_url,omitempty"`
	BlobsURL           *string          `json:"blobs_url,omitempty"`
	BranchesURL        *string          `json:"branches_url,omitempty"`
	CollaboratorsURL   *string          `json:"collaborators_url,omitempty"`
	CommentsURL        *string          `json:"comments_url,omitempty"`
	CommitsURL         *string          `json:"commits_url,omitempty"`
	CompareURL         *string          `json:"compare_url,omitempty"`
	ContentsURL        *string          `json:"contents_url,omitempty"`
	ContributorsURL    *string          `json:"contributors_url,omitempty"`
	DeploymentsURL     *string          `json:"deployments_url,omitempty"`
	DownloadsURL       *string          `json:"downloads_url,omitempty"`
	EventsURL          *string          `json:"events_url,omitempty"`
	ForksURL           *string          `json:"forks_url,omitempty"`
	GitCommitsURL      *string          `json:"git_commits_url,omitempty"`
	GitRefsURL         *string          `json:"git_refs_url,omitempty"`
	GitTagsURL         *string          `json:"git_tags_url,omitempty"`
	GitURL             *string          `json:"git_url,omitempty"`
	IssueCommentURL    *string          `json:"issue_comment_url,omitempty"`
	IssueEventsURL     *string          `json:"issue_events_url,omitempty"`
	IssuesURL          *string          `json:"issues_url,omitempty"`
	KeysURL            *string          `json:"keys_url,omitempty"`
	LabelsURL          *string          `json:"labels_url,omitempty"`
	LanguagesURL       *string          `json:"languages_url,omitempty"`
	MergesURL          *string          `json:"merges_url,omitempty"`
	MilestonesURL      *string          `json:"milestones_url,omitempty"`
	NotificationsURL   *string          `json:"notifications_url,omitempty"`
	PullsURL           *string          `json:"pulls_url,omitempty"`
	ReleasesURL        *string          `json:"releases_url,omitempty"`
	SSHURL             *string          `json:"ssh_url,omitempty"`
	StargazersURL      *string          `json:"stargazers_url,omitempty"`
	StatusesURL        *string          `json:"statuses_url,omitempty"`
	SubscribersURL     *string          `json:"subscribers_url,omitempty"`
	SubscriptionURL    *string          `json:"subscription_url,omitempty"`
	TagsURL            *string          `json:"tags_url,omitempty"`
	TeamsURL           *string          `json:"teams_url,omitempty"`
	TreesURL           *string          `json:"trees_url,omitempty"`
	CloneURL           *string          `json:"clone_url,omitempty"`
	MirrorURL          *string          `json:"mirror_url,omitempty"`
	HooksURL           *string          `json:"hooks_url,omitempty"`
	SVNURL             *string          `json:"svn_url,omitempty"`
	Homepage           *string          `json:"homepage,omitempty"`
	Language           *string          `json:"language,omitempty"`
	ForksCount         *int             `json:"forks_count,omitempty"`
	StargazersCount    *int             `json:"stargazers_count,omitempty"`
	WatchersCount      *int             `json:"watchers_count,omitempty"`
	Size               *int             `json:"size,omitempty"`
	DefaultBranch      *string          `json:"default_branch,omitempty"`
	OpenIssuesCount    *int             `json:"open_issues_count,omitempty"`
	IsTemplate         *bool            `json:"is_template,omitempty"`
	Topics             []string         `json:"topics,omitempty"`
	HasIssues          *bool            `json:"has_issues,omitempty"`
	HasProjects        *bool            `json:"has_projects,omitempty"`
	HasWiki            *bool            `json:"has_wiki,omitempty"`
	HasPages           *bool            `json:"has_pages,omitempty"`
	HasDownloads       *bool            `json:"has_downloads,omitempty"`
	Archived           *bool            `json:"archived,omitempty"`
	Disabled           *bool            `json:"disabled,omitempty"`
	PushedAt           *Timestamp       `json:"pushed_at,omitempty"`
	CreatedAt          *Timestamp       `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp       `json:"updated_at,omitempty"`
	Permissions        *map[string]bool `json:"permissions,omitempty"`
	AllowRebaseMerge   *bool            `json:"allow_rebase_merge,omitempty"`
	TemplateRepository *Repository      `json:"template_repository,omitempty"`
	AllowSquashMerge   *bool            `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit   *bool            `json:"allow_merge_commit,omitempty"`
	SubscribersCount   *int             `json:"subscribers_count,omitempty"`
	NetworkCount       *int             `json:"network_count,omitempty"`
}

func newRepository(repo *github.Repository) *Repository {
	if repo == nil {
		return nil
	}
	r := &Repository{}
	r.ID = repo.ID
	r.NodeID = repo.NodeID
	r.Name = repo.Name
	r.FullName = repo.FullName
	r.Owner = newUser(repo.Owner)
	r.Private = repo.Private
	r.HTMLURL = repo.HTMLURL
	r.Description = repo.Description
	r.Fork = repo.Fork
	r.URL = repo.URL
	r.ArchiveURL = repo.ArchiveURL
	r.AssigneesURL = repo.AssigneesURL
	r.BlobsURL = repo.BlobsURL
	r.BranchesURL = repo.BranchesURL
	r.CollaboratorsURL = repo.CollaboratorsURL
	r.CommentsURL = repo.CommentsURL
	r.CommitsURL = repo.CommitsURL
	r.CompareURL = repo.CompareURL
	r.ContentsURL = repo.ContentsURL
	r.ContributorsURL = repo.ContributorsURL
	r.DeploymentsURL = repo.DeploymentsURL
	r.DownloadsURL = repo.DownloadsURL
	r.EventsURL = repo.EventsURL
	r.ForksURL = repo.ForksURL
	r.GitCommitsURL = repo.GitCommitsURL
	r.GitRefsURL = repo.GitRefsURL
	r.GitTagsURL = repo.GitTagsURL
	r.GitURL = repo.GitURL
	r.IssueCommentURL = repo.IssueCommentURL
	r.IssueEventsURL = repo.IssueEventsURL
	r.IssuesURL = repo.IssuesURL
	r.KeysURL = repo.KeysURL
	r.LabelsURL = repo.LabelsURL
	r.LanguagesURL = repo.LanguagesURL
	r.MergesURL = repo.MergesURL
	r.MilestonesURL = repo.MilestonesURL
	r.NotificationsURL = repo.NotificationsURL
	r.PullsURL = repo.PullsURL
	r.ReleasesURL = repo.ReleasesURL
	r.SSHURL = repo.SSHURL
	r.StargazersURL = repo.StargazersURL
	r.StatusesURL = repo.StatusesURL
	r.SubscribersURL = repo.SubscribersURL
	r.SubscriptionURL = repo.SubscriptionURL
	r.TagsURL = repo.TagsURL
	r.TeamsURL = repo.TeamsURL
	r.TreesURL = repo.TreesURL
	r.CloneURL = repo.CloneURL
	r.MirrorURL = repo.MirrorURL
	r.HooksURL = repo.HooksURL
	r.SVNURL = repo.SVNURL
	r.Homepage = repo.Homepage
	r.Language = repo.Language
	r.ForksCount = repo.ForksCount
	r.StargazersCount = repo.StargazersCount
	r.WatchersCount = repo.WatchersCount
	r.Size = repo.Size
	r.DefaultBranch = repo.DefaultBranch
	r.OpenIssuesCount = repo.OpenIssuesCount
	r.IsTemplate = repo.IsTemplate
	r.Topics = repo.Topics
	r.HasIssues = repo.HasIssues
	r.HasProjects = repo.HasProjects
	r.HasWiki = repo.HasWiki
	r.HasPages = repo.HasPages
	r.HasDownloads = repo.HasDownloads
	r.Archived = repo.Archived
	r.Disabled = repo.Disabled
	r.PushedAt = newTimestamp(&repo.PushedAt.Time)
	r.CreatedAt = newTimestamp(&repo.CreatedAt.Time)
	r.UpdatedAt = newTimestamp(&repo.UpdatedAt.Time)
	r.Permissions = repo.Permissions
	r.AllowRebaseMerge = repo.AllowRebaseMerge
	r.TemplateRepository = newRepository(repo.TemplateRepository)
	r.AllowSquashMerge = repo.AllowSquashMerge
	r.AllowMergeCommit = repo.AllowMergeCommit
	r.SubscribersCount = repo.SubscribersCount
	r.NetworkCount = repo.NetworkCount
	return r
}

type PullRequestBranch struct {
	Label *string     `json:"label,omitempty"`
	Ref   *string     `json:"ref,omitempty"`
	SHA   *string     `json:"sha,omitempty"`
	User  *User       `json:"user,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
}

func newPullRequestBranch(branch *github.PullRequestBranch) *PullRequestBranch {
	if branch == nil {
		return nil
	}
	b := &PullRequestBranch{}
	b.Label = branch.Label
	b.Ref = branch.Ref
	b.SHA = branch.SHA
	b.User = newUser(branch.User)
	b.Repo = newRepository(branch.Repo)
	return b
}

type Label struct {
	ID          *int64  `json:"id,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

func newLabel(label *github.Label) *Label {
	if label == nil {
		return nil
	}
	l := &Label{}
	l.ID = label.ID
	l.NodeID = label.NodeID
	l.URL = label.URL
	l.Name = label.Name
	l.Description = label.Description
	l.Color = label.Color
	l.Default = label.Default
	return l
}

type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int64     `json:"id,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *Timestamp `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp `json:"updated_at,omitempty"`
	ClosedAt     *Timestamp `json:"closed_at,omitempty"`
	DueOn        *Timestamp `json:"due_on,omitempty"`
}

func newMilestone(milestone *github.Milestone) *Milestone {
	if milestone == nil {
		return nil
	}
	m := &Milestone{}
	m.URL = milestone.URL
	m.HTMLURL = milestone.HTMLURL
	m.LabelsURL = milestone.LabelsURL
	m.ID = milestone.ID
	m.NodeID = milestone.NodeID
	m.Number = milestone.Number
	m.State = milestone.State
	m.Title = milestone.Title
	m.Description = milestone.Description
	m.Creator = newUser(milestone.Creator)
	m.OpenIssues = milestone.OpenIssues
	m.ClosedIssues = milestone.ClosedIssues
	m.CreatedAt = newTimestamp(milestone.CreatedAt)
	m.UpdatedAt = newTimestamp(milestone.UpdatedAt)
	m.ClosedAt = newTimestamp(milestone.ClosedAt)
	m.DueOn = newTimestamp(milestone.DueOn)
	return m
}

type PullRequestComment struct {
	URL                 *string    `json:"url,omitempty"`
	ID                  *int64     `json:"id,omitempty"`
	NodeID              *string    `json:"node_id,omitempty"`
	PullRequestReviewID *int64     `json:"pull_request_review_id,omitempty"`
	DiffHunk            *string    `json:"diff_hunk,omitempty"`
	Path                *string    `json:"path,omitempty"`
	Position            *int       `json:"position,omitempty"`
	OriginalPosition    *int       `json:"original_position,omitempty"`
	CommitID            *string    `json:"commit_id,omitempty"`
	OriginalCommitID    *string    `json:"original_commit_id,omitempty"`
	InReplyTo           *int64     `json:"in_reply_to_id,omitempty"`
	User                *User      `json:"user,omitempty"`
	Body                *string    `json:"body,omitempty"`
	CreatedAt           *Timestamp `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp `json:"updated_at,omitempty"`
	HTMLURL             *string    `json:"html_url,omitempty"`
	PullRequestURL      *string    `json:"pull_request_url,omitempty"`
	AuthorAssociation   *string    `json:"author_association,omitempty"`
}

func newPullRequestComment(comment *github.PullRequestComment) *PullRequestComment {
	if comment == nil {
		return nil
	}
	c := &PullRequestComment{}
	c.URL = comment.URL
	c.ID = comment.ID
	c.NodeID = comment.NodeID
	c.PullRequestReviewID = comment.PullRequestReviewID
	c.DiffHunk = comment.DiffHunk
	c.Path = comment.Path
	c.Position = comment.Position
	c.OriginalPosition = comment.OriginalPosition
	c.CommitID = comment.CommitID
	c.OriginalCommitID = comment.OriginalCommitID
	c.InReplyTo = comment.InReplyTo
	c.User = newUser(comment.User)
	c.Body = comment.Body
	c.CreatedAt = newTimestamp(comment.CreatedAt)
	c.UpdatedAt = newTimestamp(comment.UpdatedAt)
	c.HTMLURL = comment.HTMLURL
	c.PullRequestURL = comment.PullRequestURL
	c.AuthorAssociation = comment.AuthorAssociation
	return c
}

type PullRequestReview struct {
	ID             *int64  `json:"id,omitempty"`
	NodeID         *string `json:"node_id,omitempty"`
	User           *User   `json:"user,omitempty"`
	Body           *string `json:"body,omitempty"`
	CommitID       *string `json:"commit_id,omitempty"`
	State          *string `json:"state,omitempty"`
	HTMLURL        *string `json:"html_url,omitempty"`
	PullRequestURL *string `json:"pull_request_url,omitempty"`
}

func newPullRequestReview(review *github.PullRequestReview) *PullRequestReview {
	if review == nil {
		return nil
	}
	r := &PullRequestReview{}
	r.ID = review.ID
	r.NodeID = review.NodeID
	r.User = newUser(review.User)
	r.Body = review.Body
	r.CommitID = review.CommitID
	r.State = review.State
	r.HTMLURL = review.HTMLURL
	r.PullRequestURL = review.PullRequestURL
	return r
}

type CommitAuthor struct {
	Date  *Timestamp `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
}

func newCommitAuthor(author *github.CommitAuthor) *CommitAuthor {
	if author == nil {
		return nil
	}
	ca := &CommitAuthor{}
	ca.Date = newTimestamp(author.Date)
	ca.Name = author.Name
	ca.Email = author.Email
	return ca
}

type SignatureVerification struct {
	Verified  *bool   `json:"verified,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	Signature *string `json:"signature,omitempty"`
	Payload   *string `json:"payload,omitempty"`
}

func newSignatureVerification(sig *github.SignatureVerification) *SignatureVerification {
	if sig == nil {
		return nil
	}
	s := &SignatureVerification{}
	s.Verified = sig.Verified
	s.Reason = sig.Reason
	s.Signature = sig.Signature
	s.Payload = sig.Payload
	return s
}

type Commit struct {
	URL          *string                `json:"url,omitempty"`
	Author       *CommitAuthor          `json:"author,omitempty"`
	Committer    *CommitAuthor          `json:"committer,omitempty"`
	Message      *string                `json:"message,omitempty"`
	CommentCount *int                   `json:"comment_count,omitempty"`
	Verification *SignatureVerification `json:"verification,omitempty"`
}

func newCommit(commit *github.Commit) *Commit {
	if commit == nil {
		return nil
	}
	c := &Commit{}
	c.URL = commit.URL
	c.Author = newCommitAuthor(commit.Author)
	c.Committer = newCommitAuthor(commit.Committer)
	c.Message = commit.Message
	//c.Tree = newTree(commit.Tree) // tree response specification is unknown
	c.CommentCount = commit.CommentCount
	c.Verification = newSignatureVerification(commit.Verification)
	return c
}

type RepositoryCommit struct {
	URL         *string  `json:"url,omitempty"`
	SHA         *string  `json:"sha,omitempty"`
	NodeID      *string  `json:"node_id,omitempty"`
	HTMLURL     *string  `json:"html_url,omitempty"`
	CommentsURL *string  `json:"comments_url,omitempty"`
	Commit      *Commit  `json:"commit,omitempty"`
	Author      *User    `json:"author,omitempty"`
	Committer   *User    `json:"committer,omitempty"`
	Parents     []Commit `json:"parents,omitempty"`
}

func newRepositoryCommit(commit *github.RepositoryCommit) *RepositoryCommit {
	if commit == nil {
		return nil
	}
	c := &RepositoryCommit{}
	c.URL = commit.URL
	c.SHA = commit.SHA
	c.NodeID = commit.NodeID
	c.HTMLURL = commit.HTMLURL
	c.CommentsURL = commit.CommentsURL
	c.Commit = newCommit(commit.Commit)
	c.Author = newUser(commit.Author)
	c.Committer = newUser(commit.Committer)
	c.Parents = make([]Commit, len(commit.Parents))
	for i, v := range commit.Parents {
		c.Parents[i] = *newCommit(&v)
	}
	return c
}

type RepoStatus struct {
	URL         *string    `json:"url,omitempty"`
	ID          *int64     `json:"id,omitempty"`
	NodeID      *string    `json:"node_id,omitempty"`
	State       *string    `json:"state,omitempty"`
	Description *string    `json:"description,omitempty"`
	TargetURL   *string    `json:"target_url,omitempty"`
	Context     *string    `json:"context,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	Creator     *User      `json:"creator,omitempty"`
}

func newRepoStatus(status *github.RepoStatus) *RepoStatus {
	if status == nil {
		return nil
	}
	s := &RepoStatus{}
	s.URL = status.URL
	//s.AvatarURL = status.AvatarURL // go-github not support
	s.ID = status.ID
	s.NodeID = status.NodeID
	s.State = status.State
	s.Description = status.Description
	s.TargetURL = status.TargetURL
	s.Context = status.Context
	s.CreatedAt = newTimestamp(status.CreatedAt)
	s.UpdatedAt = newTimestamp(status.UpdatedAt)
	s.Creator = newUser(status.Creator)
	return s
}

type PullRequest struct {
	URL                 *string               `json:"url,omitempty"`
	ID                  *int64                `json:"id,omitempty"`
	NodeID              *string               `json:"node_id,omitempty"`
	HTMLURL             *string               `json:"html_url,omitempty"`
	DiffURL             *string               `json:"diff_url,omitempty"`
	PatchURL            *string               `json:"patch_url,omitempty"`
	IssueURL            *string               `json:"issue_url,omitempty"`
	CommitsURL          *string               `json:"commits_url,omitempty"`
	ReviewCommentsURL   *string               `json:"review_comments_url,omitempty"`
	ReviewCommentURL    *string               `json:"review_comment_url,omitempty"`
	CommentsURL         *string               `json:"comments_url,omitempty"`
	StatusesURL         *string               `json:"statuses_url,omitempty"`
	Number              *int                  `json:"number,omitempty"`
	State               *string               `json:"state,omitempty"`
	Locked              *bool                 `json:"locked,omitempty"`
	Title               *string               `json:"title,omitempty"`
	User                *User                 `json:"user,omitempty"`
	Body                *string               `json:"body,omitempty"`
	Labels              []*Label              `json:"labels,omitempty"`
	Milestone           *Milestone            `json:"milestone,omitempty"`
	ActiveLockReason    *string               `json:"active_lock_reason,omitempty"`
	CreatedAt           *Timestamp            `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp            `json:"updated_at,omitempty"`
	ClosedAt            *Timestamp            `json:"closed_at,omitempty"`
	MergedAt            *Timestamp            `json:"merged_at,omitempty"`
	MergeCommitSHA      *string               `json:"merge_commit_sha,omitempty"`
	Assignee            *User                 `json:"assignee,omitempty"`
	Assignees           []*User               `json:"assignees,omitempty"`
	RequestedReviewers  []*User               `json:"requested_reviewers,omitempty"`
	RequestedTeams      []*Team               `json:"requested_teams,omitempty"`
	Head                *PullRequestBranch    `json:"head,omitempty"`
	Base                *PullRequestBranch    `json:"base,omitempty"`
	AuthorAssociation   *string               `json:"author_association,omitempty"`
	Draft               *bool                 `json:"draft,omitempty"`
	Merged              *bool                 `json:"merged,omitempty"`
	Mergeable           *bool                 `json:"mergeable,omitempty"`
	Rebaseable          *bool                 `json:"rebaseable,omitempty"`
	MergeableState      *string               `json:"mergeable_state,omitempty"`
	MergedBy            *User                 `json:"merged_by,omitempty"`
	ReviewComments      *int                  `json:"review_comments,omitempty"`
	MaintainerCanModify *bool                 `json:"maintainer_can_modify,omitempty"`
	Additions           *int                  `json:"additions,omitempty"`
	Deletions           *int                  `json:"deletions,omitempty"`
	ChangedFiles        *int                  `json:"changed_files,omitempty"`
	Comments            []*PullRequestComment `json:"comments"`
	Reviews             []*PullRequestReview  `json:"reviews"`
	Commits             []*RepositoryCommit   `json:"commits"`
	Statuses            []*RepoStatus         `json:"statuses"`
	Owner               *string               `json:"-"`
	Repo                *string               `json:"-"`
}

func newPullRequest(owner string, repo string, pull *github.PullRequest, comments []*github.PullRequestComment, reviews []*github.PullRequestReview, commits []*github.RepositoryCommit, statuses []*github.RepoStatus) *PullRequest {
	if pull == nil {
		return nil
	}
	p := &PullRequest{}
	p.URL = pull.URL
	p.ID = pull.ID
	p.NodeID = pull.NodeID
	p.HTMLURL = pull.HTMLURL
	p.DiffURL = pull.DiffURL
	p.PatchURL = pull.PatchURL
	p.IssueURL = pull.IssueURL
	p.CommitsURL = pull.CommitsURL
	p.ReviewCommentsURL = pull.ReviewCommentsURL
	p.ReviewCommentURL = pull.ReviewCommentURL
	p.CommentsURL = pull.CommentsURL
	p.StatusesURL = pull.StatusesURL
	p.Number = pull.Number
	p.State = pull.State
	p.Locked = pull.Locked
	p.Title = pull.Title
	p.User = newUser(pull.User)
	p.Body = pull.Body
	p.Labels = make([]*Label, len(pull.Labels))
	for i, v := range pull.Labels {
		p.Labels[i] = newLabel(v)
	}
	p.Milestone = newMilestone(pull.Milestone)
	p.ActiveLockReason = pull.ActiveLockReason
	p.CreatedAt = newTimestamp(pull.CreatedAt)
	p.UpdatedAt = newTimestamp(pull.UpdatedAt)
	p.ClosedAt = newTimestamp(pull.ClosedAt)
	p.MergedAt = newTimestamp(pull.MergedAt)
	p.MergeCommitSHA = pull.MergeCommitSHA
	p.Assignee = newUser(pull.Assignee)
	p.Assignees = make([]*User, len(pull.Assignees))
	for i, v := range pull.Assignees {
		p.Assignees[i] = newUser(v)
	}
	p.RequestedReviewers = make([]*User, len(pull.RequestedReviewers))
	for i, v := range pull.RequestedReviewers {
		p.RequestedReviewers[i] = newUser(v)
	}
	p.RequestedTeams = make([]*Team, len(pull.RequestedTeams))
	for i, v := range pull.RequestedTeams {
		p.RequestedTeams[i] = newTeam(v)
	}
	p.Head = newPullRequestBranch(pull.Head)
	p.Base = newPullRequestBranch(pull.Base)
	p.AuthorAssociation = pull.AuthorAssociation
	p.Draft = pull.Draft
	p.Merged = pull.Merged
	p.Mergeable = pull.Mergeable
	p.Rebaseable = pull.Rebaseable
	p.MergeableState = pull.MergeableState
	p.MergedBy = newUser(pull.MergedBy)
	p.ReviewComments = pull.ReviewComments
	p.MaintainerCanModify = pull.MaintainerCanModify
	p.Additions = pull.Additions
	p.Deletions = pull.Deletions
	p.ChangedFiles = pull.ChangedFiles
	p.Comments = make([]*PullRequestComment, len(comments))
	for i, comment := range comments {
		p.Comments[i] = newPullRequestComment(comment)
	}
	p.Reviews = make([]*PullRequestReview, len(reviews))
	for i, review := range reviews {
		p.Reviews[i] = newPullRequestReview(review)
	}
	p.Commits = make([]*RepositoryCommit, len(commits))
	for i, commit := range commits {
		p.Commits[i] = newRepositoryCommit(commit)
	}
	p.Statuses = make([]*RepoStatus, len(statuses))
	for i, status := range statuses {
		p.Statuses[i] = newRepoStatus(status)
	}
	p.Owner = &owner
	p.Repo = &repo
	return p
}
