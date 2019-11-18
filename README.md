# PR

[![Release](https://img.shields.io/github/v/release/k-kinzal/pr.svg?style=flat-square)](https://github.com/k-kinzal/pr/releases/latest)
[![CircleCI](https://circleci.com/gh/k-kinzal/pr.svg?style=shield)](https://circleci.com/gh/k-kinzal/pr)
[![GolangCI](https://golangci.com/badges/github.com/k-kinzal/pr.svg)](https://golangci.com/r/github.com/k-kinzal/pr)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr?ref=badge_shield)

PR is a CLI tool that operates Pull Request on a rule-based basis.

## Get Started

```bash
$ curl -L https://github.com/k-kinzal/pr/releases/download/v0.2.1/pr_linux_amd64.tar.gz | tar xz
$ cp pr /usr/local/bin/pr
$ pr --help
PR operates multiple Pull Request

Usage:
  pr [flags]
  pr [command]

Available Commands:
  check       Check if PR matches the rule and change PR status
  help        Help about any command
  merge       Merge PR that matches a rule
  show        Show PR that matches a rule
  validate    Validate the rules

Flags:
      --exit-code      returns an exit code of 127 if no PR matches the rule
  -h, --help           help for pr
      --no-exit-code   always returns 0 even if an error occurs
      --rate int       API call seconds rate limit (default 10)
      --token string   personal access token to manipulate PR [GITHUB_TOKEN]
      --version        version for pr

Use "pr [command] --help" for more information about a command.
```

## Operations

### Merge

Merge PRs that match the rule.

```bash
$ pr merge [owner]/[repo] --with-statuses -l 'state == `"open"`' -l 'length(statuses[?state == `"success"`]) > `3`'
[...]
```

## Check

When the PR CLI is run on the CI, the rule status is displayed separately from the CI.
This is a solution to the problem where multiple CI statuses are displayed in GitHub Action.

```bash
$ pr check [owner]/[repo] -l 'number == `1`' -l 'state == `"open"`' -l 'length(statuses[?state == `"success"` && context == `"ci/circleci: test"`]) == `1`'
[...]
```

Check commands can perform conditional actions.

```bash
$ pr check [owner]/[repo] --merge -l 'number == `1`' -l 'state == `"open"`' -l 'length(statuses[?state == `"success"`]) == `1`'
[...]
```

For PR with `number == 1`, merge if the condition is met, or change status to pending if the condition is not met.

## Show

Check the PR that matches the rule.

```bash
$ pr show [owner]/[repo] -l 'state == `"open"`'
[...]
```

If you want to make an error if there is no PR that matches the rule, specify `--exit-code``.

```bash
$ pr show [owner]/[repo] --exit-code -l 'number == `1`' -l 'state == `"open"`'
[...]
```

## Validate

Validate the rules.

```bash
$ pr validate [owner]/[repo] --with-statuses -l 'state == `"open"`' -l 'length(statuses[?state == `"success"`]) > `0`' -l 'user.name == `"github-action[bot]"`'
[x] state == `"open"`: 1 PRs matched the rules
[x] length(statuses[?state == `"success"`]) > `0`: 1 PRs matched the rules
[ ] user.name == `"github-action[bot]"`: no PR matches the rule
[]
```

## Rule Specification

### JSON

[JSON](https://github.com/k-kinzal/pr/blob/master/doc/spec.json)

See below for a detailed description of each item.

- [Pull Requests](https://developer.github.com/v3/pulls/)
- [Comments](https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request)
- [Reviews](https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request)
- [Commits](https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request)
- [Statuses](https://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref)
- [Checks](https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-specific-ref)

`"comments"`, `"reviews"`, `"commits"`, `"statuses""`, and `"checks"` cannot be used by default with PR link relation.
If the parameter is required, specify option `--with-comments`, `--with-reviews`, `--with-commits`, `--with-statuses`, `--with-checks` or `--with-all`.

NOTE: `--with-all` options have very poor performance. Not recommended for uses other than debugging.

### Rule Expression

In PR CLI, rules are specified using [JMESPath](http://jmespath.org/).

```bash
$ pr show [owner]/[repo] -l 'state == `"open"`' -l 'length(statuses[?state == `"success"`]) >= 1'
```

```
[?state == `"open"`] | [?length(statuses[?state == `"success"`]) >= 1])]
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
$ pr show [owner]/[repo] -l 'now() == `"2006-01-02T15:04:05Z"`'
```

`now()` returns the current unix time.

### GitHub Action

If you execute PR CLI with [GitHub Action], the rules are automatically completed by event type.

**Number completion**

- pull_request
- pull_request_review
- pull_request_review_comment

```
number == `[Pull Request Number]`
```

**Head branch completion**
- create
- deployment
- deployment_status
- push
- release

```
head.ref == `"[Branch Name]"`
```

**SHA completion**
- page_build
- status

```
head.sha == `\"[SHA]\"`
```

Please see the [event trigger](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows) for details.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr?ref=badge_large)
