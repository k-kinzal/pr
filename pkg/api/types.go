package api

import (
	"time"

	"github.com/google/go-github/v28/github"
)

type Timestamp int64

func newTimestamp(t *time.Time) Timestamp {
	if t == nil {
		return 0
	}
	return Timestamp(t.UTC().Unix())
}

type Team struct {
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Privacy         string `json:"privacy"`
	Permission      string `json:"permission"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	Parent          *Team  `json:"parent,omitempty"`
}

func newTeam(team *github.Team) *Team {
	if team == nil {
		return nil
	}
	t := &Team{}
	t.ID = team.GetID()
	t.NodeID = team.GetNodeID()
	t.URL = team.GetURL()
	//t.HTMLURL = team.GetHTMLURL() // go-github not support
	t.Name = team.GetName()
	t.Slug = team.GetSlug()
	t.Description = team.GetDescription()
	t.Privacy = team.GetPrivacy()
	t.Permission = team.GetPermission()
	t.MembersURL = team.GetMembersURL()
	t.RepositoriesURL = team.GetRepositoriesURL()
	t.Parent = newTeam(team.GetParent())
	return t
}

type User struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func newUser(user *github.User) *User {
	if user == nil {
		return nil
	}
	u := &User{}
	u.Login = user.GetLogin()
	u.ID = user.GetID()
	u.NodeID = user.GetNodeID()
	u.AvatarURL = user.GetAvatarURL()
	u.GravatarID = user.GetGravatarID()
	u.URL = user.GetURL()
	u.HTMLURL = user.GetHTMLURL()
	u.FollowersURL = user.GetFollowersURL()
	u.FollowingURL = user.GetFollowingURL()
	u.GistsURL = user.GetGistsURL()
	u.StarredURL = user.GetStarredURL()
	u.SubscriptionsURL = user.GetSubscriptionsURL()
	u.OrganizationsURL = user.GetOrganizationsURL()
	u.ReposURL = user.GetReposURL()
	u.EventsURL = user.GetEventsURL()
	u.ReceivedEventsURL = user.GetReceivedEventsURL()
	u.Type = user.GetType()
	u.SiteAdmin = user.GetSiteAdmin()
	return u
}

type Repository struct {
	ID                 int64           `json:"id"`
	NodeID             string          `json:"node_id"`
	Name               string          `json:"name"`
	FullName           string          `json:"full_name"`
	Owner              *User           `json:"owner,omitempty"`
	Private            bool            `json:"private"`
	HTMLURL            string          `json:"html_url"`
	Description        string          `json:"description"`
	Fork               bool            `json:"fork"`
	URL                string          `json:"url"`
	ArchiveURL         string          `json:"archive_url"`
	AssigneesURL       string          `json:"assignees_url"`
	BlobsURL           string          `json:"blobs_url"`
	BranchesURL        string          `json:"branches_url"`
	CollaboratorsURL   string          `json:"collaborators_url"`
	CommentsURL        string          `json:"comments_url"`
	CommitsURL         string          `json:"commits_url"`
	CompareURL         string          `json:"compare_url"`
	ContentsURL        string          `json:"contents_url"`
	ContributorsURL    string          `json:"contributors_url"`
	DeploymentsURL     string          `json:"deployments_url"`
	DownloadsURL       string          `json:"downloads_url"`
	EventsURL          string          `json:"events_url"`
	ForksURL           string          `json:"forks_url"`
	GitCommitsURL      string          `json:"git_commits_url"`
	GitRefsURL         string          `json:"git_refs_url"`
	GitTagsURL         string          `json:"git_tags_url"`
	GitURL             string          `json:"git_url"`
	IssueCommentURL    string          `json:"issue_comment_url"`
	IssueEventsURL     string          `json:"issue_events_url"`
	IssuesURL          string          `json:"issues_url"`
	KeysURL            string          `json:"keys_url"`
	LabelsURL          string          `json:"labels_url"`
	LanguagesURL       string          `json:"languages_url"`
	MergesURL          string          `json:"merges_url"`
	MilestonesURL      string          `json:"milestones_url"`
	NotificationsURL   string          `json:"notifications_url"`
	PullsURL           string          `json:"pulls_url"`
	ReleasesURL        string          `json:"releases_url"`
	SSHURL             string          `json:"ssh_url"`
	StargazersURL      string          `json:"stargazers_url"`
	StatusesURL        string          `json:"statuses_url"`
	SubscribersURL     string          `json:"subscribers_url"`
	SubscriptionURL    string          `json:"subscription_url"`
	TagsURL            string          `json:"tags_url"`
	TeamsURL           string          `json:"teams_url"`
	TreesURL           string          `json:"trees_url"`
	CloneURL           string          `json:"clone_url"`
	MirrorURL          string          `json:"mirror_url"`
	HooksURL           string          `json:"hooks_url"`
	SVNURL             string          `json:"svn_url"`
	Homepage           string          `json:"homepage"`
	Language           string          `json:"language"`
	ForksCount         int             `json:"forks_count"`
	StargazersCount    int             `json:"stargazers_count"`
	WatchersCount      int             `json:"watchers_count"`
	Size               int             `json:"size"`
	DefaultBranch      string          `json:"default_branch"`
	OpenIssuesCount    int             `json:"open_issues_count"`
	IsTemplate         bool            `json:"is_template"`
	HasIssues          bool            `json:"has_issues"`
	HasProjects        bool            `json:"has_projects"`
	HasWiki            bool            `json:"has_wiki"`
	HasPages           bool            `json:"has_pages"`
	HasDownloads       bool            `json:"has_downloads"`
	Archived           bool            `json:"archived"`
	Disabled           bool            `json:"disabled"`
	PushedAt           Timestamp       `json:"pushed_at,omitempty"`
	CreatedAt          Timestamp       `json:"created_at,omitempty"`
	UpdatedAt          Timestamp       `json:"updated_at,omitempty"`
	Permissions        map[string]bool `json:"permissions"`
	AllowRebaseMerge   bool            `json:"allow_rebase_merge"`
	TemplateRepository *Repository     `json:"template_repository,omitempty"`
	AllowSquashMerge   bool            `json:"allow_squash_merge"`
	AllowMergeCommit   bool            `json:"allow_merge_commit"`
	SubscribersCount   int             `json:"subscribers_count"`
	NetworkCount       int             `json:"network_count"`
}

func newRepository(repo *github.Repository) *Repository {
	if repo == nil {
		return nil
	}
	r := &Repository{}
	r.ID = repo.GetID()
	r.NodeID = repo.GetNodeID()
	r.Name = repo.GetName()
	r.FullName = repo.GetFullName()
	r.Owner = newUser(repo.GetOwner())
	r.Private = repo.GetPrivate()
	r.HTMLURL = repo.GetHTMLURL()
	r.Description = repo.GetDescription()
	r.Fork = repo.GetFork()
	r.URL = repo.GetURL()
	r.ArchiveURL = repo.GetArchiveURL()
	r.AssigneesURL = repo.GetAssigneesURL()
	r.BlobsURL = repo.GetBlobsURL()
	r.BranchesURL = repo.GetBranchesURL()
	r.CollaboratorsURL = repo.GetCollaboratorsURL()
	r.CommentsURL = repo.GetCommentsURL()
	r.CommitsURL = repo.GetCommitsURL()
	r.CompareURL = repo.GetCompareURL()
	r.ContentsURL = repo.GetContentsURL()
	r.ContributorsURL = repo.GetContributorsURL()
	r.DeploymentsURL = repo.GetDeploymentsURL()
	r.DownloadsURL = repo.GetDownloadsURL()
	r.EventsURL = repo.GetEventsURL()
	r.ForksURL = repo.GetForksURL()
	r.GitCommitsURL = repo.GetGitCommitsURL()
	r.GitRefsURL = repo.GetGitRefsURL()
	r.GitTagsURL = repo.GetGitTagsURL()
	r.GitURL = repo.GetGitURL()
	r.IssueCommentURL = repo.GetIssueCommentURL()
	r.IssueEventsURL = repo.GetIssueEventsURL()
	r.IssuesURL = repo.GetIssuesURL()
	r.KeysURL = repo.GetKeysURL()
	r.LabelsURL = repo.GetLabelsURL()
	r.LanguagesURL = repo.GetLanguagesURL()
	r.MergesURL = repo.GetMergesURL()
	r.MilestonesURL = repo.GetMilestonesURL()
	r.NotificationsURL = repo.GetNotificationsURL()
	r.PullsURL = repo.GetPullsURL()
	r.ReleasesURL = repo.GetReleasesURL()
	r.SSHURL = repo.GetSSHURL()
	r.StargazersURL = repo.GetStargazersURL()
	r.StatusesURL = repo.GetStatusesURL()
	r.SubscribersURL = repo.GetSubscribersURL()
	r.SubscriptionURL = repo.GetSubscriptionURL()
	r.TagsURL = repo.GetTagsURL()
	r.TeamsURL = repo.GetTeamsURL()
	r.TreesURL = repo.GetTreesURL()
	r.CloneURL = repo.GetCloneURL()
	r.MirrorURL = repo.GetMirrorURL()
	r.HooksURL = repo.GetHooksURL()
	r.SVNURL = repo.GetSVNURL()
	r.Homepage = repo.GetHomepage()
	r.Language = repo.GetLanguage()
	r.ForksCount = repo.GetForksCount()
	r.StargazersCount = repo.GetStargazersCount()
	r.WatchersCount = repo.GetWatchersCount()
	r.Size = repo.GetSize()
	r.DefaultBranch = repo.GetDefaultBranch()
	r.OpenIssuesCount = repo.GetOpenIssuesCount()
	r.IsTemplate = repo.GetIsTemplate()
	//r.Topics = repo.GetTopics() // go-github not support
	r.HasIssues = repo.GetHasIssues()
	r.HasProjects = repo.GetHasProjects()
	r.HasWiki = repo.GetHasWiki()
	r.HasPages = repo.GetHasPages()
	r.HasDownloads = repo.GetHasDownloads()
	r.Archived = repo.GetArchived()
	r.Disabled = repo.GetDisabled()
	t1 := repo.GetPushedAt().Time
	r.PushedAt = newTimestamp(&t1)
	t2 := repo.GetCreatedAt().Time
	r.CreatedAt = newTimestamp(&t2)
	t3 := repo.GetUpdatedAt().Time
	r.UpdatedAt = newTimestamp(&t3)
	r.Permissions = repo.GetPermissions()
	r.AllowRebaseMerge = repo.GetAllowRebaseMerge()
	r.TemplateRepository = newRepository(repo.GetTemplateRepository())
	r.AllowSquashMerge = repo.GetAllowSquashMerge()
	r.AllowMergeCommit = repo.GetAllowMergeCommit()
	r.SubscribersCount = repo.GetSubscribersCount()
	r.NetworkCount = repo.GetNetworkCount()
	return r
}

type PullRequestBranch struct {
	Label string      `json:"label"`
	Ref   string      `json:"ref"`
	SHA   string      `json:"sha"`
	User  *User       `json:"user,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
}

func newPullRequestBranch(branch *github.PullRequestBranch) *PullRequestBranch {
	if branch == nil {
		return nil
	}
	b := &PullRequestBranch{}
	b.Label = branch.GetLabel()
	b.Ref = branch.GetRef()
	b.SHA = branch.GetSHA()
	b.User = newUser(branch.GetUser())
	b.Repo = newRepository(branch.GetRepo())
	return b
}

type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

func newLabel(label *github.Label) *Label {
	if label == nil {
		return nil
	}
	l := &Label{}
	l.ID = label.GetID()
	l.NodeID = label.GetNodeID()
	l.URL = label.GetURL()
	l.Name = label.GetName()
	l.Description = label.GetDescription()
	l.Color = label.GetColor()
	l.Default = label.GetDefault()
	return l
}

type Milestone struct {
	URL          string    `json:"url"`
	HTMLURL      string    `json:"html_url"`
	LabelsURL    string    `json:"labels_url"`
	ID           int64     `json:"id"`
	NodeID       string    `json:"node_id"`
	Number       int       `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Creator      *User     `json:"creator,omitempty"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    Timestamp `json:"created_at,omitempty"`
	UpdatedAt    Timestamp `json:"updated_at,omitempty"`
	ClosedAt     Timestamp `json:"closed_at,omitempty"`
	DueOn        Timestamp `json:"due_on,omitempty"`
}

func newMilestone(milestone *github.Milestone) *Milestone {
	if milestone == nil {
		return nil
	}
	m := &Milestone{}
	m.URL = milestone.GetURL()
	m.HTMLURL = milestone.GetHTMLURL()
	m.LabelsURL = milestone.GetLabelsURL()
	m.ID = milestone.GetID()
	m.NodeID = milestone.GetNodeID()
	m.Number = milestone.GetNumber()
	m.State = milestone.GetState()
	m.Title = milestone.GetTitle()
	m.Description = milestone.GetDescription()
	m.Creator = newUser(milestone.GetCreator())
	m.OpenIssues = milestone.GetOpenIssues()
	m.ClosedIssues = milestone.GetClosedIssues()
	m.CreatedAt = newTimestamp(milestone.CreatedAt)
	m.UpdatedAt = newTimestamp(milestone.UpdatedAt)
	m.ClosedAt = newTimestamp(milestone.ClosedAt)
	m.DueOn = newTimestamp(milestone.DueOn)
	return m
}

type PullRequestComment struct {
	URL                 string    `json:"url"`
	ID                  int64     `json:"id"`
	NodeID              string    `json:"node_id"`
	PullRequestReviewID int64     `json:"pull_request_review_id"`
	DiffHunk            string    `json:"diff_hunk"`
	Path                string    `json:"path"`
	Position            int       `json:"position"`
	OriginalPosition    int       `json:"original_position"`
	CommitID            string    `json:"commit_id"`
	OriginalCommitID    string    `json:"original_commit_id"`
	InReplyTo           int64     `json:"in_reply_to_id"`
	User                *User     `json:"user,omitempty"`
	Body                string    `json:"body"`
	CreatedAt           Timestamp `json:"created_at,omitempty"`
	UpdatedAt           Timestamp `json:"updated_at,omitempty"`
	HTMLURL             string    `json:"html_url"`
	PullRequestURL      string    `json:"pull_request_url"`
	AuthorAssociation   string    `json:"author_association"`
}

func newPullRequestComment(comment *github.PullRequestComment) *PullRequestComment {
	if comment == nil {
		return nil
	}
	c := &PullRequestComment{}
	c.URL = comment.GetURL()
	c.ID = comment.GetID()
	c.NodeID = comment.GetNodeID()
	c.PullRequestReviewID = comment.GetPullRequestReviewID()
	c.DiffHunk = comment.GetDiffHunk()
	c.Path = comment.GetPath()
	c.Position = comment.GetPosition()
	c.OriginalPosition = comment.GetOriginalPosition()
	c.CommitID = comment.GetCommitID()
	c.OriginalCommitID = comment.GetOriginalCommitID()
	c.InReplyTo = comment.GetInReplyTo()
	c.User = newUser(comment.GetUser())
	c.Body = comment.GetBody()
	c.CreatedAt = newTimestamp(comment.CreatedAt)
	c.UpdatedAt = newTimestamp(comment.UpdatedAt)
	c.HTMLURL = comment.GetHTMLURL()
	c.PullRequestURL = comment.GetPullRequestURL()
	c.AuthorAssociation = comment.GetAuthorAssociation()
	return c
}

type PullRequestReview struct {
	ID             int64  `json:"id"`
	NodeID         string `json:"node_id"`
	User           *User  `json:"user,omitempty"`
	Body           string `json:"body"`
	CommitID       string `json:"commit_id"`
	State          string `json:"state"`
	HTMLURL        string `json:"html_url"`
	PullRequestURL string `json:"pull_request_url"`
}

func newPullRequestReview(review *github.PullRequestReview) *PullRequestReview {
	if review == nil {
		return nil
	}
	r := &PullRequestReview{}
	r.ID = review.GetID()
	r.NodeID = review.GetNodeID()
	r.User = newUser(review.GetUser())
	r.Body = review.GetBody()
	r.CommitID = review.GetCommitID()
	r.State = review.GetState()
	r.HTMLURL = review.GetHTMLURL()
	r.PullRequestURL = review.GetPullRequestURL()
	return r
}

type CommitAuthor struct {
	Date  Timestamp `json:"date"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func newCommitAuthor(author *github.CommitAuthor) *CommitAuthor {
	if author == nil {
		return nil
	}
	ca := &CommitAuthor{}
	ca.Date = newTimestamp(author.Date)
	ca.Name = author.GetName()
	ca.Email = author.GetEmail()
	return ca
}

type SignatureVerification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

func newSignatureVerification(sig *github.SignatureVerification) *SignatureVerification {
	if sig == nil {
		return nil
	}
	s := &SignatureVerification{}
	s.Verified = sig.GetVerified()
	s.Reason = sig.GetReason()
	s.Signature = sig.GetSignature()
	s.Payload = sig.GetPayload()
	return s
}

type Commit struct {
	URL          string                 `json:"url"`
	Author       *CommitAuthor          `json:"author,omitempty"`
	Committer    *CommitAuthor          `json:"committer,omitempty"`
	Message      string                 `json:"message"`
	CommentCount int                    `json:"comment_count"`
	Verification *SignatureVerification `json:"verification,omitempty"`
}

func newCommit(commit *github.Commit) *Commit {
	if commit == nil {
		return nil
	}
	c := &Commit{}
	c.URL = commit.GetURL()
	c.Author = newCommitAuthor(commit.GetAuthor())
	c.Committer = newCommitAuthor(commit.GetCommitter())
	c.Message = commit.GetMessage()
	//c.Tree = newTree(commit.GetTree()) // tree response specification is unknown
	c.CommentCount = commit.GetCommentCount()
	c.Verification = newSignatureVerification(commit.GetVerification())
	return c
}

type RepositoryCommit struct {
	URL         string    `json:"url"`
	SHA         string    `json:"sha"`
	NodeID      string    `json:"node_id"`
	HTMLURL     string    `json:"html_url"`
	CommentsURL string    `json:"comments_url"`
	Commit      *Commit   `json:"commit,omitempty"`
	Author      *User     `json:"author,omitempty"`
	Committer   *User     `json:"committer,omitempty"`
	Parents     []*Commit `json:"parents,omitempty"`
}

func newRepositoryCommit(commit *github.RepositoryCommit) *RepositoryCommit {
	if commit == nil {
		return nil
	}
	c := &RepositoryCommit{}
	c.URL = commit.GetURL()
	c.SHA = commit.GetSHA()
	c.NodeID = commit.GetNodeID()
	c.HTMLURL = commit.GetHTMLURL()
	c.CommentsURL = commit.GetCommentsURL()
	c.Commit = newCommit(commit.GetCommit())
	c.Author = newUser(commit.GetAuthor())
	c.Committer = newUser(commit.GetCommitter())
	c.Parents = make([]*Commit, len(commit.Parents))
	for i, v := range commit.Parents {
		c.Parents[i] = newCommit(&v)
	}
	return c
}

type RepoStatus struct {
	URL         string    `json:"url"`
	ID          int64     `json:"id"`
	NodeID      string    `json:"node_id"`
	State       string    `json:"state"`
	Description string    `json:"description"`
	TargetURL   string    `json:"target_url"`
	Context     string    `json:"context"`
	CreatedAt   Timestamp `json:"created_at,omitempty"`
	UpdatedAt   Timestamp `json:"updated_at,omitempty"`
	Creator     *User     `json:"creator,omitempty"`
}

func newRepoStatus(status *github.RepoStatus) *RepoStatus {
	if status == nil {
		return nil
	}
	s := &RepoStatus{}
	s.URL = status.GetURL()
	//s.AvatarURL = status.GetAvatarURL // go-github not support
	s.ID = status.GetID()
	s.NodeID = status.GetNodeID()
	s.State = status.GetState()
	s.Description = status.GetDescription()
	s.TargetURL = status.GetTargetURL()
	s.Context = status.GetContext()
	s.CreatedAt = newTimestamp(status.CreatedAt)
	s.UpdatedAt = newTimestamp(status.UpdatedAt)
	s.Creator = newUser(status.GetCreator())
	return s
}

type PullRequest struct {
	URL                 string                `json:"url"`
	ID                  int64                 `json:"id"`
	NodeID              string                `json:"node_id"`
	HTMLURL             string                `json:"html_url"`
	DiffURL             string                `json:"diff_url"`
	PatchURL            string                `json:"patch_url"`
	IssueURL            string                `json:"issue_url"`
	CommitsURL          string                `json:"commits_url"`
	ReviewCommentsURL   string                `json:"review_comments_url"`
	ReviewCommentURL    string                `json:"review_comment_url"`
	CommentsURL         string                `json:"comments_url"`
	StatusesURL         string                `json:"statuses_url"`
	Number              int                   `json:"number"`
	State               string                `json:"state"`
	Locked              bool                  `json:"locked"`
	Title               string                `json:"title"`
	User                *User                 `json:"user,omitempty"`
	Body                string                `json:"body"`
	Labels              []*Label              `json:"labels,omitempty"`
	Milestone           *Milestone            `json:"milestone,omitempty"`
	ActiveLockReason    string                `json:"active_lock_reason"`
	CreatedAt           Timestamp             `json:"created_at,omitempty"`
	UpdatedAt           Timestamp             `json:"updated_at,omitempty"`
	ClosedAt            Timestamp             `json:"closed_at,omitempty"`
	MergedAt            Timestamp             `json:"merged_at,omitempty"`
	MergeCommitSHA      string                `json:"merge_commit_sha"`
	Assignee            *User                 `json:"assignee,omitempty"`
	Assignees           []*User               `json:"assignees,omitempty"`
	RequestedReviewers  []*User               `json:"requested_reviewers,omitempty"`
	RequestedTeams      []*Team               `json:"requested_teams,omitempty"`
	Head                *PullRequestBranch    `json:"head,omitempty"`
	Base                *PullRequestBranch    `json:"base,omitempty"`
	AuthorAssociation   string                `json:"author_association"`
	Draft               bool                  `json:"draft"`
	Merged              bool                  `json:"merged"`
	Mergeable           bool                  `json:"mergeable"`
	Rebaseable          bool                  `json:"rebaseable"`
	MergeableState      string                `json:"mergeable_state"`
	MergedBy            *User                 `json:"merged_by,omitempty"`
	ReviewComments      int                   `json:"review_comments"`
	MaintainerCanModify bool                  `json:"maintainer_can_modify"`
	Additions           int                   `json:"additions"`
	Deletions           int                   `json:"deletions"`
	ChangedFiles        int                   `json:"changed_files"`
	Comments            []*PullRequestComment `json:"comments"`
	Reviews             []*PullRequestReview  `json:"reviews"`
	Commits             []*RepositoryCommit   `json:"commits"`
	Statuses            []*RepoStatus         `json:"statuses"`
	Owner               string                `json:"-"`
	Repo                string                `json:"-"`
}

func newPullRequest(owner string, repo string, pull *github.PullRequest, comments []*github.PullRequestComment, reviews []*github.PullRequestReview, commits []*github.RepositoryCommit, statuses []*github.RepoStatus) *PullRequest {
	if pull == nil {
		return nil
	}
	p := &PullRequest{}
	p.URL = pull.GetURL()
	p.ID = pull.GetID()
	p.NodeID = pull.GetNodeID()
	p.HTMLURL = pull.GetHTMLURL()
	p.DiffURL = pull.GetDiffURL()
	p.PatchURL = pull.GetPatchURL()
	p.IssueURL = pull.GetIssueURL()
	p.CommitsURL = pull.GetCommitsURL()
	p.ReviewCommentsURL = pull.GetReviewCommentsURL()
	p.ReviewCommentURL = pull.GetReviewCommentURL()
	p.CommentsURL = pull.GetCommentsURL()
	p.StatusesURL = pull.GetStatusesURL()
	p.Number = pull.GetNumber()
	p.State = pull.GetState()
	p.Locked = pull.GetLocked()
	p.Title = pull.GetTitle()
	p.User = newUser(pull.GetUser())
	p.Body = pull.GetBody()
	p.Labels = make([]*Label, len(pull.Labels))
	for i, v := range pull.Labels {
		p.Labels[i] = newLabel(v)
	}
	p.Milestone = newMilestone(pull.GetMilestone())
	p.ActiveLockReason = pull.GetActiveLockReason()
	p.CreatedAt = newTimestamp(pull.CreatedAt)
	p.UpdatedAt = newTimestamp(pull.UpdatedAt)
	p.ClosedAt = newTimestamp(pull.ClosedAt)
	p.MergedAt = newTimestamp(pull.MergedAt)
	p.MergeCommitSHA = pull.GetMergeCommitSHA()
	p.Assignee = newUser(pull.GetAssignee())
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
	p.Head = newPullRequestBranch(pull.GetHead())
	p.Base = newPullRequestBranch(pull.GetBase())
	p.AuthorAssociation = pull.GetAuthorAssociation()
	p.Draft = pull.GetDraft()
	p.Merged = pull.GetMerged()
	p.Mergeable = pull.GetMergeable()
	p.Rebaseable = pull.GetRebaseable()
	p.MergeableState = pull.GetMergeableState()
	p.MergedBy = newUser(pull.GetMergedBy())
	p.ReviewComments = pull.GetReviewComments()
	p.MaintainerCanModify = pull.GetMaintainerCanModify()
	p.Additions = pull.GetAdditions()
	p.Deletions = pull.GetDeletions()
	p.ChangedFiles = pull.GetChangedFiles()
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
	p.Owner = owner
	p.Repo = repo
	return p
}
