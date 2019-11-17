package api

import (
	"time"

	"github.com/google/go-github/v28/github"
)

type Timestamp float64

func newTimestamp(t *time.Time) Timestamp {
	if t == nil {
		return 0
	}
	return Timestamp(t.UTC().Unix())
}

type Team struct {
	Id              float64 `json:"id"`
	NodeId          string  `json:"node_id"`
	Url             string  `json:"url"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	Description     string  `json:"description"`
	Privacy         string  `json:"privacy"`
	Permission      string  `json:"permission"`
	MembersUrl      string  `json:"members_url"`
	RepositoriesUrl string  `json:"repositories_url"`
	Parent          *Team   `json:"parent,omitempty"`
}

func newTeam(team *github.Team) *Team {
	if team == nil {
		return nil
	}
	t := &Team{}
	t.Id = float64(team.GetID())
	t.NodeId = team.GetNodeID()
	t.Url = team.GetURL()
	//t.HTMLURL = team.GetHTMLURL() // go-github not support
	t.Name = team.GetName()
	t.Slug = team.GetSlug()
	t.Description = team.GetDescription()
	t.Privacy = team.GetPrivacy()
	t.Permission = team.GetPermission()
	t.MembersUrl = team.GetMembersURL()
	t.RepositoriesUrl = team.GetRepositoriesURL()
	t.Parent = newTeam(team.GetParent())
	return t
}

type User struct {
	Login             string  `json:"login"`
	Id                float64 `json:"id"`
	NodeId            string  `json:"node_id"`
	AvatarUrl         string  `json:"avatar_url"`
	GravatarId        string  `json:"gravatar_id"`
	Url               string  `json:"url"`
	HtmlUrl           string  `json:"html_url"`
	FollowersUrl      string  `json:"followers_url"`
	FollowingUrl      string  `json:"following_url"`
	GistsUrl          string  `json:"gists_url"`
	StarredUrl        string  `json:"starred_url"`
	SubscriptionsUrl  string  `json:"subscriptions_url"`
	OrganizationsUrl  string  `json:"organizations_url"`
	ReposUrl          string  `json:"repos_url"`
	EventsUrl         string  `json:"events_url"`
	ReceivedEventsUrl string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
}

func newUser(user *github.User) *User {
	if user == nil {
		return nil
	}
	u := &User{}
	u.Login = user.GetLogin()
	u.Id = float64(user.GetID())
	u.NodeId = user.GetNodeID()
	u.AvatarUrl = user.GetAvatarURL()
	u.GravatarId = user.GetGravatarID()
	u.Url = user.GetURL()
	u.HtmlUrl = user.GetHTMLURL()
	u.FollowersUrl = user.GetFollowersURL()
	u.FollowingUrl = user.GetFollowingURL()
	u.GistsUrl = user.GetGistsURL()
	u.StarredUrl = user.GetStarredURL()
	u.SubscriptionsUrl = user.GetSubscriptionsURL()
	u.OrganizationsUrl = user.GetOrganizationsURL()
	u.ReposUrl = user.GetReposURL()
	u.EventsUrl = user.GetEventsURL()
	u.ReceivedEventsUrl = user.GetReceivedEventsURL()
	u.Type = user.GetType()
	u.SiteAdmin = user.GetSiteAdmin()
	return u
}

type Repository struct {
	Id                 float64         `json:"id"`
	NodeId             string          `json:"node_id"`
	Name               string          `json:"name"`
	FullName           string          `json:"full_name"`
	Owner              *User           `json:"owner,omitempty"`
	Private            bool            `json:"private"`
	HtmlUrl            string          `json:"html_url"`
	Description        string          `json:"description"`
	Fork               bool            `json:"fork"`
	Url                string          `json:"url"`
	ArchiveUrl         string          `json:"archive_url"`
	AssigneesUrl       string          `json:"assignees_url"`
	BlobsUrl           string          `json:"blobs_url"`
	BranchesUrl        string          `json:"branches_url"`
	CollaboratorsUrl   string          `json:"collaborators_url"`
	CommentsUrl        string          `json:"comments_url"`
	CommitsUrl         string          `json:"commits_url"`
	CompareUrl         string          `json:"compare_url"`
	ContentsUrl        string          `json:"contents_url"`
	ContributorsUrl    string          `json:"contributors_url"`
	DeploymentsUrl     string          `json:"deployments_url"`
	DownloadsUrl       string          `json:"downloads_url"`
	EventsUrl          string          `json:"events_url"`
	ForksUrl           string          `json:"forks_url"`
	GitCommitsUrl      string          `json:"git_commits_url"`
	GitRefsUrl         string          `json:"git_refs_url"`
	GitTagsUrl         string          `json:"git_tags_url"`
	GitUrl             string          `json:"git_url"`
	IssueCommentUrl    string          `json:"issue_comment_url"`
	IssueEventsUrl     string          `json:"issue_events_url"`
	IssuesUrl          string          `json:"issues_url"`
	KeysUrl            string          `json:"keys_url"`
	LabelsUrl          string          `json:"labels_url"`
	LanguagesUrl       string          `json:"languages_url"`
	MergesUrl          string          `json:"merges_url"`
	MilestonesUrl      string          `json:"milestones_url"`
	NotificationsUrl   string          `json:"notifications_url"`
	PullsUrl           string          `json:"pulls_url"`
	ReleasesUrl        string          `json:"releases_url"`
	SshUrl             string          `json:"ssh_url"`
	StargazersUrl      string          `json:"stargazers_url"`
	StatusesUrl        string          `json:"statuses_url"`
	SubscribersUrl     string          `json:"subscribers_url"`
	SubscriptionUrl    string          `json:"subscription_url"`
	TagsUrl            string          `json:"tags_url"`
	TeamsUrl           string          `json:"teams_url"`
	TreesUrl           string          `json:"trees_url"`
	CloneUrl           string          `json:"clone_url"`
	MirrorUrl          string          `json:"mirror_url"`
	HooksUrl           string          `json:"hooks_url"`
	SvnUrl             string          `json:"svn_url"`
	Homepage           string          `json:"homepage"`
	Language           string          `json:"language"`
	ForksCount         float64         `json:"forks_count"`
	StargazersCount    float64         `json:"stargazers_count"`
	WatchersCount      float64         `json:"watchers_count"`
	Size               float64         `json:"size"`
	DefaultBranch      string          `json:"default_branch"`
	OpenIssuesCount    float64         `json:"open_issues_count"`
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
	SubscribersCount   float64         `json:"subscribers_count"`
	NetworkCount       float64         `json:"network_count"`
}

func newRepository(repo *github.Repository) *Repository {
	if repo == nil {
		return nil
	}
	r := &Repository{}
	r.Id = float64(repo.GetID())
	r.NodeId = repo.GetNodeID()
	r.Name = repo.GetName()
	r.FullName = repo.GetFullName()
	r.Owner = newUser(repo.GetOwner())
	r.Private = repo.GetPrivate()
	r.HtmlUrl = repo.GetHTMLURL()
	r.Description = repo.GetDescription()
	r.Fork = repo.GetFork()
	r.Url = repo.GetURL()
	r.ArchiveUrl = repo.GetArchiveURL()
	r.AssigneesUrl = repo.GetAssigneesURL()
	r.BlobsUrl = repo.GetBlobsURL()
	r.BranchesUrl = repo.GetBranchesURL()
	r.CollaboratorsUrl = repo.GetCollaboratorsURL()
	r.CommentsUrl = repo.GetCommentsURL()
	r.CommitsUrl = repo.GetCommitsURL()
	r.CompareUrl = repo.GetCompareURL()
	r.ContentsUrl = repo.GetContentsURL()
	r.ContributorsUrl = repo.GetContributorsURL()
	r.DeploymentsUrl = repo.GetDeploymentsURL()
	r.DownloadsUrl = repo.GetDownloadsURL()
	r.EventsUrl = repo.GetEventsURL()
	r.ForksUrl = repo.GetForksURL()
	r.GitCommitsUrl = repo.GetGitCommitsURL()
	r.GitRefsUrl = repo.GetGitRefsURL()
	r.GitTagsUrl = repo.GetGitTagsURL()
	r.GitUrl = repo.GetGitURL()
	r.IssueCommentUrl = repo.GetIssueCommentURL()
	r.IssueEventsUrl = repo.GetIssueEventsURL()
	r.IssuesUrl = repo.GetIssuesURL()
	r.KeysUrl = repo.GetKeysURL()
	r.LabelsUrl = repo.GetLabelsURL()
	r.LanguagesUrl = repo.GetLanguagesURL()
	r.MergesUrl = repo.GetMergesURL()
	r.MilestonesUrl = repo.GetMilestonesURL()
	r.NotificationsUrl = repo.GetNotificationsURL()
	r.PullsUrl = repo.GetPullsURL()
	r.ReleasesUrl = repo.GetReleasesURL()
	r.SshUrl = repo.GetSSHURL()
	r.StargazersUrl = repo.GetStargazersURL()
	r.StatusesUrl = repo.GetStatusesURL()
	r.SubscribersUrl = repo.GetSubscribersURL()
	r.SubscriptionUrl = repo.GetSubscriptionURL()
	r.TagsUrl = repo.GetTagsURL()
	r.TeamsUrl = repo.GetTeamsURL()
	r.TreesUrl = repo.GetTreesURL()
	r.CloneUrl = repo.GetCloneURL()
	r.MirrorUrl = repo.GetMirrorURL()
	r.HooksUrl = repo.GetHooksURL()
	r.SvnUrl = repo.GetSVNURL()
	r.Homepage = repo.GetHomepage()
	r.Language = repo.GetLanguage()
	r.ForksCount = float64(repo.GetForksCount())
	r.StargazersCount = float64(repo.GetStargazersCount())
	r.WatchersCount = float64(repo.GetWatchersCount())
	r.Size = float64(repo.GetSize())
	r.DefaultBranch = repo.GetDefaultBranch()
	r.OpenIssuesCount = float64(repo.GetOpenIssuesCount())
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
	r.SubscribersCount = float64(repo.GetSubscribersCount())
	r.NetworkCount = float64(repo.GetNetworkCount())
	return r
}

type PullRequestBranch struct {
	Label string      `json:"label"`
	Ref   string      `json:"ref"`
	Sha   string      `json:"sha"`
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
	b.Sha = branch.GetSHA()
	b.User = newUser(branch.GetUser())
	b.Repo = newRepository(branch.GetRepo())
	return b
}

type Label struct {
	Id          float64 `json:"id"`
	NodeId      string  `json:"node_id"`
	Url         string  `json:"url"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	Default     bool    `json:"default"`
}

func newLabel(label *github.Label) *Label {
	if label == nil {
		return nil
	}
	l := &Label{}
	l.Id = float64(label.GetID())
	l.NodeId = label.GetNodeID()
	l.Url = label.GetURL()
	l.Name = label.GetName()
	l.Description = label.GetDescription()
	l.Color = label.GetColor()
	l.Default = label.GetDefault()
	return l
}

type Milestone struct {
	Url          string    `json:"url"`
	HtmlUrl      string    `json:"html_url"`
	LabelsUrl    string    `json:"labels_url"`
	Id           float64   `json:"id"`
	NodeId       string    `json:"node_id"`
	Number       float64   `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Creator      *User     `json:"creator,omitempty"`
	OpenIssues   float64   `json:"open_issues"`
	ClosedIssues float64   `json:"closed_issues"`
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
	m.Url = milestone.GetURL()
	m.HtmlUrl = milestone.GetHTMLURL()
	m.LabelsUrl = milestone.GetLabelsURL()
	m.Id = float64(milestone.GetID())
	m.NodeId = milestone.GetNodeID()
	m.Number = float64(milestone.GetNumber())
	m.State = milestone.GetState()
	m.Title = milestone.GetTitle()
	m.Description = milestone.GetDescription()
	m.Creator = newUser(milestone.GetCreator())
	m.OpenIssues = float64(milestone.GetOpenIssues())
	m.ClosedIssues = float64(milestone.GetClosedIssues())
	m.CreatedAt = newTimestamp(milestone.CreatedAt)
	m.UpdatedAt = newTimestamp(milestone.UpdatedAt)
	m.ClosedAt = newTimestamp(milestone.ClosedAt)
	m.DueOn = newTimestamp(milestone.DueOn)
	return m
}

type PullRequestComment struct {
	Url                 string    `json:"url"`
	Id                  float64   `json:"id"`
	NodeId              string    `json:"node_id"`
	PullRequestReviewId float64   `json:"pull_request_review_id"`
	DiffHunk            string    `json:"diff_hunk"`
	Path                string    `json:"path"`
	Position            float64   `json:"position"`
	OriginalPosition    float64   `json:"original_position"`
	CommitId            string    `json:"commit_id"`
	OriginalCommitId    string    `json:"original_commit_id"`
	InReplyTo           float64   `json:"in_reply_to_id"`
	User                *User     `json:"user,omitempty"`
	Body                string    `json:"body"`
	CreatedAt           Timestamp `json:"created_at,omitempty"`
	UpdatedAt           Timestamp `json:"updated_at,omitempty"`
	HtmlUrl             string    `json:"html_url"`
	PullRequestUrl      string    `json:"pull_request_url"`
	AuthorAssociation   string    `json:"author_association"`
}

func newPullRequestComment(comment *github.PullRequestComment) *PullRequestComment {
	if comment == nil {
		return nil
	}
	c := &PullRequestComment{}
	c.Url = comment.GetURL()
	c.Id = float64(comment.GetID())
	c.NodeId = comment.GetNodeID()
	c.PullRequestReviewId = float64(comment.GetPullRequestReviewID())
	c.DiffHunk = comment.GetDiffHunk()
	c.Path = comment.GetPath()
	c.Position = float64(comment.GetPosition())
	c.OriginalPosition = float64(comment.GetOriginalPosition())
	c.CommitId = comment.GetCommitID()
	c.OriginalCommitId = comment.GetOriginalCommitID()
	c.InReplyTo = float64(comment.GetInReplyTo())
	c.User = newUser(comment.GetUser())
	c.Body = comment.GetBody()
	c.CreatedAt = newTimestamp(comment.CreatedAt)
	c.UpdatedAt = newTimestamp(comment.UpdatedAt)
	c.HtmlUrl = comment.GetHTMLURL()
	c.PullRequestUrl = comment.GetPullRequestURL()
	c.AuthorAssociation = comment.GetAuthorAssociation()
	return c
}

type PullRequestReview struct {
	Id             float64 `json:"id"`
	NodeId         string  `json:"node_id"`
	User           *User   `json:"user,omitempty"`
	Body           string  `json:"body"`
	CommitId       string  `json:"commit_id"`
	State          string  `json:"state"`
	HtmlUrl        string  `json:"html_url"`
	PullRequestUrl string  `json:"pull_request_url"`
}

func newPullRequestReview(review *github.PullRequestReview) *PullRequestReview {
	if review == nil {
		return nil
	}
	r := &PullRequestReview{}
	r.Id = float64(review.GetID())
	r.NodeId = review.GetNodeID()
	r.User = newUser(review.GetUser())
	r.Body = review.GetBody()
	r.CommitId = review.GetCommitID()
	r.State = review.GetState()
	r.HtmlUrl = review.GetHTMLURL()
	r.PullRequestUrl = review.GetPullRequestURL()
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
	Url          string                 `json:"url"`
	Author       *CommitAuthor          `json:"author,omitempty"`
	Committer    *CommitAuthor          `json:"committer,omitempty"`
	Message      string                 `json:"message"`
	CommentCount float64                `json:"comment_count"`
	Verification *SignatureVerification `json:"verification,omitempty"`
}

func newCommit(commit *github.Commit) *Commit {
	if commit == nil {
		return nil
	}
	c := &Commit{}
	c.Url = commit.GetURL()
	c.Author = newCommitAuthor(commit.GetAuthor())
	c.Committer = newCommitAuthor(commit.GetCommitter())
	c.Message = commit.GetMessage()
	//c.Tree = newTree(commit.GetTree()) // tree response specification is unknown
	c.CommentCount = float64(commit.GetCommentCount())
	c.Verification = newSignatureVerification(commit.GetVerification())
	return c
}

type RepositoryCommit struct {
	Url         string    `json:"url"`
	Sha         string    `json:"sha"`
	NodeId      string    `json:"node_id"`
	HtmlUrl     string    `json:"html_url"`
	CommentsUrl string    `json:"comments_url"`
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
	c.Url = commit.GetURL()
	c.Sha = commit.GetSHA()
	c.NodeId = commit.GetNodeID()
	c.HtmlUrl = commit.GetHTMLURL()
	c.CommentsUrl = commit.GetCommentsURL()
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
	Url         string    `json:"url"`
	Id          float64   `json:"id"`
	NodeId      string    `json:"node_id"`
	State       string    `json:"state"`
	Description string    `json:"description"`
	TargetUrl   string    `json:"target_url"`
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
	s.Url = status.GetURL()
	//s.AvatarUrl = status.GetAvatarURL() // go-github not support
	s.Id = float64(status.GetID())
	s.NodeId = status.GetNodeID()
	s.State = status.GetState()
	s.Description = status.GetDescription()
	s.TargetUrl = status.GetTargetURL()
	s.Context = status.GetContext()
	s.CreatedAt = newTimestamp(status.CreatedAt)
	s.UpdatedAt = newTimestamp(status.UpdatedAt)
	s.Creator = newUser(status.GetCreator())
	return s
}

type App struct {
	Id          float64   `json:"id,omitempty"`
	NodeId      string    `json:"node_id,omitempty"`
	Owner       *User     `json:"owner,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	ExternalUrl string    `json:"external_url,omitempty"`
	HtmlUrl     string    `json:"html_url,omitempty"`
	CreatedAt   Timestamp `json:"created_at,omitempty"`
	UpdatedAt   Timestamp `json:"updated_at,omitempty"`
}

func newApp(app *github.App) *App {
	if app == nil {
		return nil
	}
	a := &App{}
	a.Id = float64(app.GetID())
	a.NodeId = app.GetNodeID()
	a.Owner = newUser(app.GetOwner())
	a.Name = app.GetName()
	a.Description = app.GetDescription()
	a.ExternalUrl = app.GetExternalURL()
	a.HtmlUrl = app.GetHTMLURL()
	t1 := app.GetCreatedAt().Time
	a.CreatedAt = newTimestamp(&t1)
	t2 := app.GetUpdatedAt().Time
	a.UpdatedAt = newTimestamp(&t2)
	return a
}

type CheckSuite struct {
	Id         float64     `json:"id,omitempty"`
	NodeId     string      `json:"node_id,omitempty"`
	HeadBranch string      `json:"head_branch,omitempty"`
	HeadSha    string      `json:"head_sha,omitempty"`
	Url        string      `json:"url,omitempty"`
	BeforeSha  string      `json:"before,omitempty"`
	AfterSha   string      `json:"after,omitempty"`
	Status     string      `json:"status,omitempty"`
	Conclusion string      `json:"conclusion,omitempty"`
	App        *App        `json:"app,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	HeadCommit *Commit     `json:"head_commit,omitempty"`
}

func newCheckSuite(suite *github.CheckSuite) *CheckSuite {
	if suite == nil {
		return nil
	}
	s := &CheckSuite{}
	s.Id = float64(suite.GetID())
	s.NodeId = suite.GetNodeID()
	s.HeadBranch = suite.GetHeadBranch()
	s.HeadSha = suite.GetHeadSHA()
	s.Url = suite.GetURL()
	s.BeforeSha = suite.GetBeforeSHA()
	s.AfterSha = suite.GetAfterSHA()
	s.Status = suite.GetStatus()
	s.Conclusion = suite.GetConclusion()
	s.App = newApp(suite.GetApp())
	s.Repository = newRepository(suite.GetRepository())
	s.HeadCommit = newCommit(suite.GetHeadCommit())
	return s
}

type CheckRunAnnotation struct {
	Path            string  `json:"path,omitempty"`
	BlobHRef        string  `json:"blob_href,omitempty"`
	StartLine       float64 `json:"start_line,omitempty"`
	EndLine         float64 `json:"end_line,omitempty"`
	StartColumn     float64 `json:"start_column,omitempty"`
	EndColumn       float64 `json:"end_column,omitempty"`
	AnnotationLevel string  `json:"annotation_level,omitempty"`
	Message         string  `json:"message,omitempty"`
	Title           string  `json:"title,omitempty"`
	RawDetails      string  `json:"raw_details,omitempty"`
}

func newCheckRunAnnotation(anno *github.CheckRunAnnotation) *CheckRunAnnotation {
	if anno == nil {
		return nil
	}
	a := &CheckRunAnnotation{}
	a.Path = anno.GetPath()
	a.BlobHRef = anno.GetBlobHRef()
	a.StartLine = float64(anno.GetStartLine())
	a.EndLine = float64(anno.GetEndLine())
	a.StartColumn = float64(anno.GetStartColumn())
	a.EndColumn = float64(anno.GetEndColumn())
	a.AnnotationLevel = anno.GetAnnotationLevel()
	a.Message = anno.GetMessage()
	a.Title = anno.GetTitle()
	a.RawDetails = anno.GetRawDetails()
	return a
}

type CheckRunImage struct {
	Alt      string `json:"alt,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

func newCheckRunImage(image *github.CheckRunImage) *CheckRunImage {
	if image == nil {
		return nil
	}
	i := &CheckRunImage{}
	i.Alt = image.GetAlt()
	i.ImageUrl = image.GetImageURL()
	i.Caption = image.GetCaption()
	return i
}

type CheckRunOutput struct {
	Title            string                `json:"title,omitempty"`
	Summary          string                `json:"summary,omitempty"`
	Text             string                `json:"text,omitempty"`
	AnnotationsCount float64               `json:"annotations_count,omitempty"`
	AnnotationsUrl   string                `json:"annotations_url,omitempty"`
	Annotations      []*CheckRunAnnotation `json:"annotations,omitempty"`
	Images           []*CheckRunImage      `json:"images,omitempty"`
}

func newCheckRunOutput(output *github.CheckRunOutput) *CheckRunOutput {
	if output == nil {
		return nil
	}
	o := &CheckRunOutput{}
	o.Title = output.GetTitle()
	o.Summary = output.GetSummary()
	o.Text = output.GetText()
	o.AnnotationsCount = float64(output.GetAnnotationsCount())
	o.AnnotationsUrl = output.GetAnnotationsURL()
	for _, annotation := range output.Annotations {
		o.Annotations = append(o.Annotations, newCheckRunAnnotation(annotation))
	}
	for _, image := range output.Images {
		o.Images = append(o.Images, newCheckRunImage(image))
	}
	return o
}

type CheckRun struct {
	Id          float64         `json:"id,omitempty"`
	NodeId      string          `json:"node_id,omitempty"`
	HeadSha     string          `json:"head_sha,omitempty"`
	ExternalId  string          `json:"external_id,omitempty"`
	Url         string          `json:"url,omitempty"`
	HtmlUrl     string          `json:"html_url,omitempty"`
	DetailsUrl  string          `json:"details_url,omitempty"`
	Status      string          `json:"status,omitempty"`
	Conclusion  string          `json:"conclusion,omitempty"`
	StartedAt   Timestamp       `json:"started_at,omitempty"`
	CompletedAt Timestamp       `json:"completed_at,omitempty"`
	Output      *CheckRunOutput `json:"output,omitempty"`
	Name        string          `json:"name,omitempty"`
	CheckSuite  *CheckSuite     `json:"check_suite,omitempty"`
	App         *App            `json:"app,omitempty"`
}

func newCheckRun(run *github.CheckRun) *CheckRun {
	if run == nil {
		return nil
	}
	r := &CheckRun{}
	r.Id = float64(run.GetID())
	r.NodeId = run.GetNodeID()
	r.HeadSha = run.GetHeadSHA()
	r.ExternalId = run.GetExternalID()
	r.Url = run.GetURL()
	r.HtmlUrl = run.GetHTMLURL()
	r.DetailsUrl = run.GetDetailsURL()
	r.Status = run.GetStatus()
	r.Conclusion = run.GetConclusion()
	t1 := run.GetStartedAt().Time
	r.StartedAt = newTimestamp(&t1)
	t2 := run.GetCompletedAt().Time
	r.CompletedAt = newTimestamp(&t2)
	r.Output = newCheckRunOutput(run.GetOutput())
	r.Name = run.GetName()
	r.CheckSuite = newCheckSuite(run.GetCheckSuite())
	r.App = newApp(run.GetApp())
	return r
}

type PullRequest struct {
	Url                 string                `json:"url"`
	Id                  float64               `json:"id"`
	NodeId              string                `json:"node_id"`
	HtmlUrl             string                `json:"html_url"`
	DiffUrl             string                `json:"diff_url"`
	PatchUrl            string                `json:"patch_url"`
	IssueUrl            string                `json:"issue_url"`
	CommitsUrl          string                `json:"commits_url"`
	ReviewCommentsUrl   string                `json:"review_comments_url"`
	ReviewCommentUrl    string                `json:"review_comment_url"`
	CommentsUrl         string                `json:"comments_url"`
	StatusesUrl         string                `json:"statuses_url"`
	Number              float64               `json:"number"`
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
	MergeCommitSha      string                `json:"merge_commit_sha"`
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
	ReviewComments      float64               `json:"review_comments"`
	MaintainerCanModify bool                  `json:"maintainer_can_modify"`
	Additions           float64               `json:"additions"`
	Deletions           float64               `json:"deletions"`
	ChangedFiles        float64               `json:"changed_files"`
	Comments            []*PullRequestComment `json:"comments"`
	Reviews             []*PullRequestReview  `json:"reviews"`
	Commits             []*RepositoryCommit   `json:"commits"`
	Statuses            []*RepoStatus         `json:"statuses"`
	Checks              []*CheckRun           `json:"checks"`
	Owner               string                `json:"-"`
	Repo                string                `json:"-"`
}

func newPullRequest(owner string, repo string, pull *github.PullRequest, comments []*github.PullRequestComment, reviews []*github.PullRequestReview, commits []*github.RepositoryCommit, statuses []*github.RepoStatus, checks []*github.CheckRun) *PullRequest {
	if pull == nil {
		return nil
	}
	p := &PullRequest{}
	p.Url = pull.GetURL()
	p.Id = float64(pull.GetID())
	p.NodeId = pull.GetNodeID()
	p.HtmlUrl = pull.GetHTMLURL()
	p.DiffUrl = pull.GetDiffURL()
	p.PatchUrl = pull.GetPatchURL()
	p.IssueUrl = pull.GetIssueURL()
	p.CommitsUrl = pull.GetCommitsURL()
	p.ReviewCommentsUrl = pull.GetReviewCommentsURL()
	p.ReviewCommentUrl = pull.GetReviewCommentURL()
	p.CommentsUrl = pull.GetCommentsURL()
	p.StatusesUrl = pull.GetStatusesURL()
	p.Number = float64(pull.GetNumber())
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
	p.MergeCommitSha = pull.GetMergeCommitSHA()
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
	p.ReviewComments = float64(pull.GetReviewComments())
	p.MaintainerCanModify = pull.GetMaintainerCanModify()
	p.Additions = float64(pull.GetAdditions())
	p.Deletions = float64(pull.GetDeletions())
	p.ChangedFiles = float64(pull.GetChangedFiles())
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
	p.Checks = make([]*CheckRun, len(checks))
	for i, check := range checks {
		p.Checks[i] = newCheckRun(check)
	}
	p.Owner = owner
	p.Repo = repo
	return p
}
