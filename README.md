# PR

[![Release](https://img.shields.io/github/v/release/k-kinzal/pr.svg?style=flat-square)](https://github.com/k-kinzal/pr/releases/latest)
[![CircleCI](https://circleci.com/gh/k-kinzal/pr.svg?style=shield)](https://circleci.com/gh/k-kinzal/pr)
[![GolangCI](https://golangci.com/badges/github.com/k-kinzal/pr.svg)](https://golangci.com/r/github.com/k-kinzal/pr)

PR is a CLI tool that operates Pull Request on a rule-based basis.

## Get Started

```bash

```

## Operations

### Merge

Merge PRs that match the rule.

```bash
$ pr merge [owner]/[repo] -l 'state == `"open"`' -l 'length(statuses) == length(statuses[?state == `"success"`])'
[...]
```

## Show

Check the PR that matches the rule.

```bash
$ pr show [owner]/[repo] -l 'state == `"open"`'
[...]
```

If you want to make an error if there is no PR that matches the rule, specify `--exit-code``.

```bash
$ pr show [owner]/[repo] --exit-code -l 'number == `1`' -l 'state == `"open"`'
[]
```

## Rule Specification

### JSON

```json
[
  {
    "url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
    "id": 1,
    "node_id": "MDExOlB1bGxSZXF1ZXN0MQ==",
    "html_url": "https://github.com/octocat/Hello-World/pull/1347",
    "diff_url": "https://github.com/octocat/Hello-World/pull/1347.diff",
    "patch_url": "https://github.com/octocat/Hello-World/pull/1347.patch",
    "issue_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347",
    "commits_url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347/commits",
    "review_comments_url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347/comments",
    "review_comment_url": "https://api.github.com/repos/octocat/Hello-World/pulls/comments{/number}",
    "comments_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments",
    "statuses_url": "https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e",
    "number": 1347,
    "state": "open",
    "locked": true,
    "title": "Amazing new feature",
    "user": {
      "login": "octocat",
      "id": 1,
      "node_id": "MDQ6VXNlcjE=",
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "repos_url": "https://api.github.com/users/octocat/repos",
      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "body": "Please pull these awesome changes in!",
    "labels": [
      {
        "id": 208045946,
        "node_id": "MDU6TGFiZWwyMDgwNDU5NDY=",
        "url": "https://api.github.com/repos/octocat/Hello-World/labels/bug",
        "name": "bug",
        "description": "Something isn't working",
        "color": "f29513",
        "default": true
      }
    ],
    "milestone": {
      "url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
      "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
      "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
      "id": 1002604,
      "node_id": "MDk6TWlsZXN0b25lMTAwMjYwNA==",
      "number": 1,
      "state": "open",
      "title": "v1.0",
      "description": "Tracking milestone for version 1.0",
      "creator": {
        "login": "octocat",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      },
      "open_issues": 4,
      "closed_issues": 8,
      "created_at": 1136214245,
      "updated_at": 1136214245,
      "closed_at": 1136214245,
      "due_on": 1136214245
    },
    "active_lock_reason": "too heated",
    "created_at": 1136214245,
    "updated_at": 1136214245,
    "closed_at": 1136214245,
    "merged_at": 1136214245,
    "merge_commit_sha": "e5bd3914e2e596debea16f433f57875b5b90bcd6",
    "assignee": {
      "login": "octocat",
      "id": 1,
      "node_id": "MDQ6VXNlcjE=",
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "repos_url": "https://api.github.com/users/octocat/repos",
      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "assignees": [
      {
        "login": "octocat",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      },
      {
        "login": "hubot",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/hubot_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/hubot",
        "html_url": "https://github.com/hubot",
        "followers_url": "https://api.github.com/users/hubot/followers",
        "following_url": "https://api.github.com/users/hubot/following{/other_user}",
        "gists_url": "https://api.github.com/users/hubot/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/hubot/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/hubot/subscriptions",
        "organizations_url": "https://api.github.com/users/hubot/orgs",
        "repos_url": "https://api.github.com/users/hubot/repos",
        "events_url": "https://api.github.com/users/hubot/events{/privacy}",
        "received_events_url": "https://api.github.com/users/hubot/received_events",
        "type": "User",
        "site_admin": true
      }
    ],
    "requested_reviewers": [
      {
        "login": "other_user",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/other_user_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/other_user",
        "html_url": "https://github.com/other_user",
        "followers_url": "https://api.github.com/users/other_user/followers",
        "following_url": "https://api.github.com/users/other_user/following{/other_user}",
        "gists_url": "https://api.github.com/users/other_user/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/other_user/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/other_user/subscriptions",
        "organizations_url": "https://api.github.com/users/other_user/orgs",
        "repos_url": "https://api.github.com/users/other_user/repos",
        "events_url": "https://api.github.com/users/other_user/events{/privacy}",
        "received_events_url": "https://api.github.com/users/other_user/received_events",
        "type": "User",
        "site_admin": false
      }
    ],
    "requested_teams": [
      {
        "id": 1,
        "node_id": "MDQ6VGVhbTE=",
        "url": "https://api.github.com/teams/1",
        "html_url": "https://api.github.com/teams/justice-league",
        "name": "Justice League",
        "slug": "justice-league",
        "description": "A great team.",
        "privacy": "closed",
        "permission": "admin",
        "members_url": "https://api.github.com/teams/1/members{/member}",
        "repositories_url": "https://api.github.com/teams/1/repos",
        "parent": null
      }
    ],
    "head": {
      "label": "octocat:new-topic",
      "ref": "new-topic",
      "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e",
      "user": {
        "login": "octocat",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 1296269,
        "node_id": "MDEwOlJlcG9zaXRvcnkxMjk2MjY5",
        "name": "Hello-World",
        "full_name": "octocat/Hello-World",
        "owner": {
          "login": "octocat",
          "id": 1,
          "node_id": "MDQ6VXNlcjE=",
          "avatar_url": "https://github.com/images/error/octocat_happy.gif",
          "gravatar_id": "",
          "url": "https://api.github.com/users/octocat",
          "html_url": "https://github.com/octocat",
          "followers_url": "https://api.github.com/users/octocat/followers",
          "following_url": "https://api.github.com/users/octocat/following{/other_user}",
          "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
          "organizations_url": "https://api.github.com/users/octocat/orgs",
          "repos_url": "https://api.github.com/users/octocat/repos",
          "events_url": "https://api.github.com/users/octocat/events{/privacy}",
          "received_events_url": "https://api.github.com/users/octocat/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/octocat/Hello-World",
        "description": "This your first repo!",
        "fork": false,
        "url": "https://api.github.com/repos/octocat/Hello-World",
        "archive_url": "http://api.github.com/repos/octocat/Hello-World/{archive_format}{/ref}",
        "assignees_url": "http://api.github.com/repos/octocat/Hello-World/assignees{/user}",
        "blobs_url": "http://api.github.com/repos/octocat/Hello-World/git/blobs{/sha}",
        "branches_url": "http://api.github.com/repos/octocat/Hello-World/branches{/branch}",
        "collaborators_url": "http://api.github.com/repos/octocat/Hello-World/collaborators{/collaborator}",
        "comments_url": "http://api.github.com/repos/octocat/Hello-World/comments{/number}",
        "commits_url": "http://api.github.com/repos/octocat/Hello-World/commits{/sha}",
        "compare_url": "http://api.github.com/repos/octocat/Hello-World/compare/{base}...{head}",
        "contents_url": "http://api.github.com/repos/octocat/Hello-World/contents/{+path}",
        "contributors_url": "http://api.github.com/repos/octocat/Hello-World/contributors",
        "deployments_url": "http://api.github.com/repos/octocat/Hello-World/deployments",
        "downloads_url": "http://api.github.com/repos/octocat/Hello-World/downloads",
        "events_url": "http://api.github.com/repos/octocat/Hello-World/events",
        "forks_url": "http://api.github.com/repos/octocat/Hello-World/forks",
        "git_commits_url": "http://api.github.com/repos/octocat/Hello-World/git/commits{/sha}",
        "git_refs_url": "http://api.github.com/repos/octocat/Hello-World/git/refs{/sha}",
        "git_tags_url": "http://api.github.com/repos/octocat/Hello-World/git/tags{/sha}",
        "git_url": "git:github.com/octocat/Hello-World.git",
        "issue_comment_url": "http://api.github.com/repos/octocat/Hello-World/issues/comments{/number}",
        "issue_events_url": "http://api.github.com/repos/octocat/Hello-World/issues/events{/number}",
        "issues_url": "http://api.github.com/repos/octocat/Hello-World/issues{/number}",
        "keys_url": "http://api.github.com/repos/octocat/Hello-World/keys{/key_id}",
        "labels_url": "http://api.github.com/repos/octocat/Hello-World/labels{/name}",
        "languages_url": "http://api.github.com/repos/octocat/Hello-World/languages",
        "merges_url": "http://api.github.com/repos/octocat/Hello-World/merges",
        "milestones_url": "http://api.github.com/repos/octocat/Hello-World/milestones{/number}",
        "notifications_url": "http://api.github.com/repos/octocat/Hello-World/notifications{?since,all,participating}",
        "pulls_url": "http://api.github.com/repos/octocat/Hello-World/pulls{/number}",
        "releases_url": "http://api.github.com/repos/octocat/Hello-World/releases{/id}",
        "ssh_url": "git@github.com:octocat/Hello-World.git",
        "stargazers_url": "http://api.github.com/repos/octocat/Hello-World/stargazers",
        "statuses_url": "http://api.github.com/repos/octocat/Hello-World/statuses/{sha}",
        "subscribers_url": "http://api.github.com/repos/octocat/Hello-World/subscribers",
        "subscription_url": "http://api.github.com/repos/octocat/Hello-World/subscription",
        "tags_url": "http://api.github.com/repos/octocat/Hello-World/tags",
        "teams_url": "http://api.github.com/repos/octocat/Hello-World/teams",
        "trees_url": "http://api.github.com/repos/octocat/Hello-World/git/trees{/sha}",
        "clone_url": "https://github.com/octocat/Hello-World.git",
        "mirror_url": "git:git.example.com/octocat/Hello-World",
        "hooks_url": "http://api.github.com/repos/octocat/Hello-World/hooks",
        "svn_url": "https://svn.github.com/octocat/Hello-World",
        "homepage": "https://github.com",
        "language": null,
        "forks_count": 9,
        "stargazers_count": 80,
        "watchers_count": 80,
        "size": 108,
        "default_branch": "master",
        "open_issues_count": 0,
        "is_template": true,
        "topics": [
          "octocat",
          "atom",
          "electron",
          "api"
        ],
        "has_issues": true,
        "has_projects": true,
        "has_wiki": true,
        "has_pages": false,
        "has_downloads": true,
        "archived": false,
        "disabled": false,
        "pushed_at": 1136214245,
        "created_at": 1136214245,
        "updated_at": 1136214245,
        "permissions": {
          "admin": false,
          "push": false,
          "pull": true
        },
        "allow_rebase_merge": true,
        "template_repository": null,
        "allow_squash_merge": true,
        "allow_merge_commit": true,
        "subscribers_count": 42,
        "network_count": 0
      }
    },
    "base": {
      "label": "octocat:master",
      "ref": "master",
      "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e",
      "user": {
        "login": "octocat",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 1296269,
        "node_id": "MDEwOlJlcG9zaXRvcnkxMjk2MjY5",
        "name": "Hello-World",
        "full_name": "octocat/Hello-World",
        "owner": {
          "login": "octocat",
          "id": 1,
          "node_id": "MDQ6VXNlcjE=",
          "avatar_url": "https://github.com/images/error/octocat_happy.gif",
          "gravatar_id": "",
          "url": "https://api.github.com/users/octocat",
          "html_url": "https://github.com/octocat",
          "followers_url": "https://api.github.com/users/octocat/followers",
          "following_url": "https://api.github.com/users/octocat/following{/other_user}",
          "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
          "organizations_url": "https://api.github.com/users/octocat/orgs",
          "repos_url": "https://api.github.com/users/octocat/repos",
          "events_url": "https://api.github.com/users/octocat/events{/privacy}",
          "received_events_url": "https://api.github.com/users/octocat/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/octocat/Hello-World",
        "description": "This your first repo!",
        "fork": false,
        "url": "https://api.github.com/repos/octocat/Hello-World",
        "archive_url": "http://api.github.com/repos/octocat/Hello-World/{archive_format}{/ref}",
        "assignees_url": "http://api.github.com/repos/octocat/Hello-World/assignees{/user}",
        "blobs_url": "http://api.github.com/repos/octocat/Hello-World/git/blobs{/sha}",
        "branches_url": "http://api.github.com/repos/octocat/Hello-World/branches{/branch}",
        "collaborators_url": "http://api.github.com/repos/octocat/Hello-World/collaborators{/collaborator}",
        "comments_url": "http://api.github.com/repos/octocat/Hello-World/comments{/number}",
        "commits_url": "http://api.github.com/repos/octocat/Hello-World/commits{/sha}",
        "compare_url": "http://api.github.com/repos/octocat/Hello-World/compare/{base}...{head}",
        "contents_url": "http://api.github.com/repos/octocat/Hello-World/contents/{+path}",
        "contributors_url": "http://api.github.com/repos/octocat/Hello-World/contributors",
        "deployments_url": "http://api.github.com/repos/octocat/Hello-World/deployments",
        "downloads_url": "http://api.github.com/repos/octocat/Hello-World/downloads",
        "events_url": "http://api.github.com/repos/octocat/Hello-World/events",
        "forks_url": "http://api.github.com/repos/octocat/Hello-World/forks",
        "git_commits_url": "http://api.github.com/repos/octocat/Hello-World/git/commits{/sha}",
        "git_refs_url": "http://api.github.com/repos/octocat/Hello-World/git/refs{/sha}",
        "git_tags_url": "http://api.github.com/repos/octocat/Hello-World/git/tags{/sha}",
        "git_url": "git:github.com/octocat/Hello-World.git",
        "issue_comment_url": "http://api.github.com/repos/octocat/Hello-World/issues/comments{/number}",
        "issue_events_url": "http://api.github.com/repos/octocat/Hello-World/issues/events{/number}",
        "issues_url": "http://api.github.com/repos/octocat/Hello-World/issues{/number}",
        "keys_url": "http://api.github.com/repos/octocat/Hello-World/keys{/key_id}",
        "labels_url": "http://api.github.com/repos/octocat/Hello-World/labels{/name}",
        "languages_url": "http://api.github.com/repos/octocat/Hello-World/languages",
        "merges_url": "http://api.github.com/repos/octocat/Hello-World/merges",
        "milestones_url": "http://api.github.com/repos/octocat/Hello-World/milestones{/number}",
        "notifications_url": "http://api.github.com/repos/octocat/Hello-World/notifications{?since,all,participating}",
        "pulls_url": "http://api.github.com/repos/octocat/Hello-World/pulls{/number}",
        "releases_url": "http://api.github.com/repos/octocat/Hello-World/releases{/id}",
        "ssh_url": "git@github.com:octocat/Hello-World.git",
        "stargazers_url": "http://api.github.com/repos/octocat/Hello-World/stargazers",
        "statuses_url": "http://api.github.com/repos/octocat/Hello-World/statuses/{sha}",
        "subscribers_url": "http://api.github.com/repos/octocat/Hello-World/subscribers",
        "subscription_url": "http://api.github.com/repos/octocat/Hello-World/subscription",
        "tags_url": "http://api.github.com/repos/octocat/Hello-World/tags",
        "teams_url": "http://api.github.com/repos/octocat/Hello-World/teams",
        "trees_url": "http://api.github.com/repos/octocat/Hello-World/git/trees{/sha}",
        "clone_url": "https://github.com/octocat/Hello-World.git",
        "mirror_url": "git:git.example.com/octocat/Hello-World",
        "hooks_url": "http://api.github.com/repos/octocat/Hello-World/hooks",
        "svn_url": "https://svn.github.com/octocat/Hello-World",
        "homepage": "https://github.com",
        "language": null,
        "forks_count": 9,
        "stargazers_count": 80,
        "watchers_count": 80,
        "size": 108,
        "default_branch": "master",
        "open_issues_count": 0,
        "is_template": true,
        "topics": [
          "octocat",
          "atom",
          "electron",
          "api"
        ],
        "has_issues": true,
        "has_projects": true,
        "has_wiki": true,
        "has_pages": false,
        "has_downloads": true,
        "archived": false,
        "disabled": false,
        "pushed_at": 1136214245,
        "created_at": 1136214245,
        "updated_at": 1136214245,
        "permissions": {
          "admin": false,
          "push": false,
          "pull": true
        },
        "allow_rebase_merge": true,
        "template_repository": null,
        "allow_squash_merge": true,
        "allow_merge_commit": true,
        "subscribers_count": 42,
        "network_count": 0
      }
    },
    "_links": {
      "self": {
        "href": "https://api.github.com/repos/octocat/Hello-World/pulls/1347"
      },
      "html": {
        "href": "https://github.com/octocat/Hello-World/pull/1347"
      },
      "issue": {
        "href": "https://api.github.com/repos/octocat/Hello-World/issues/1347"
      },
      "comments": {
        "href": "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments"
      },
      "review_comments": {
        "href": "https://api.github.com/repos/octocat/Hello-World/pulls/1347/comments"
      },
      "review_comment": {
        "href": "https://api.github.com/repos/octocat/Hello-World/pulls/comments{/number}"
      },
      "commits": {
        "href": "https://api.github.com/repos/octocat/Hello-World/pulls/1347/commits"
      },
      "statuses": {
        "href": "https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e"
      }
    },
    "author_association": "OWNER",
    "draft": false,
    "statuses": [
      {
        "url": "https://api.github.com/repos/octocat/Hello-World/statuses/6dcb09b5b57875f334f61aebed695e2e4193db5e",
        "avatar_url": "https://github.com/images/error/hubot_happy.gif",
        "id": 1,
        "node_id": "MDY6U3RhdHVzMQ==",
        "state": "success",
        "description": "Build has completed successfully",
        "target_url": "https://ci.example.com/1000/output",
        "context": "continuous-integration/jenkins",
        "created_at": 1136214245,
        "updated_at": 1136214245,
        "creator": {
          "login": "octocat",
          "id": 1,
          "node_id": "MDQ6VXNlcjE=",
          "avatar_url": "https://github.com/images/error/octocat_happy.gif",
          "gravatar_id": "",
          "url": "https://api.github.com/users/octocat",
          "html_url": "https://github.com/octocat",
          "followers_url": "https://api.github.com/users/octocat/followers",
          "following_url": "https://api.github.com/users/octocat/following{/other_user}",
          "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
          "organizations_url": "https://api.github.com/users/octocat/orgs",
          "repos_url": "https://api.github.com/users/octocat/repos",
          "events_url": "https://api.github.com/users/octocat/events{/privacy}",
          "received_events_url": "https://api.github.com/users/octocat/received_events",
          "type": "User",
          "site_admin": false
        }
      }
    ]
  }
]
```

See below for a detailed description of each item.

- [Pull Requests](https://developer.github.com/v3/pulls/)
- [Statuses](https://developer.github.com/v3/repos/statuses/)

### Rule Expression

In pr, rules are specified using [JMESPath](http://jmespath.org/).

```bash
$ pr show [owner]/[repo] -l 'state == `"open"`' -l 'length(statuses) == length(statuses[?state == `"success"`])'
```

```
[?state == `"open"`] | [?length(statuses) == length(statuses[?state == `"success"`])]
```

The specified rule is converted to an expression that combines [Filter Expression](http://jmespath.org/proposals/filter-expressions.html) with a pipe.

#### Extend date string 

The date string has been extended to be replaced with unix time.

```bash
$ pr show [owner]/[repo] -l 'now() == `"2006-01-02T15:04:05Z"`' -l 'now() > `"15:04:05"`'
```
```
[?`1571475658` >= `1136214245`] | [?`1571475658` >= `1571497445`]
```

If the date is in the format `"2006-01-02T15:04:05Z "`, it will be treated as unix time.
The format of `"15:04:05"` is regarded as time and treated as unix time for the specified time of the day.

### Extend Function

In pr, JMESPath can be extended to use original functions.

#### now()

```
$ pr show [owner]/[repo] -l
```

`now()` returns the current unix time.