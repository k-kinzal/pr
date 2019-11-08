# PR

[![Release](https://img.shields.io/github/v/release/k-kinzal/pr.svg?style=flat-square)](https://github.com/k-kinzal/pr/releases/latest)
[![CircleCI](https://circleci.com/gh/k-kinzal/pr.svg?style=shield)](https://circleci.com/gh/k-kinzal/pr)
[![GolangCI](https://golangci.com/badges/github.com/k-kinzal/pr.svg)](https://golangci.com/r/github.com/k-kinzal/pr)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr?ref=badge_shield)

PR is a CLI tool that operates Pull Request on a rule-based basis.

## Get Started

```bash
$ curl -L https://github.com/k-kinzal/pr/releases/download/v0.1.0/pr_linux_amd64.tar.gz | tar xz
$ cp pr /usr/local/bin/pr
$ pr --help
PR operates multiple Pull Request

Usage:
  pr [flags]
  pr [command]

Available Commands:
  help        Help about any command
  merge       Merge PR that matches a rule
  show        Show PR that matches a rule

Flags:
      --exit-code      returns an exit code of 127 if no PR matches the rule
  -h, --help           help for pr
      --token string   personal access token to manipulate PR [GITHUB_TOKEN]

Use "pr [command] --help" for more information about a command.
```

## Operations

### Merge

Merge PRs that match the rule.

```bash
$ pr merge [owner]/[repo] --with-statuses -l 'state == `"open"`' -l 'length(statuses) == length(statuses[?state == `"success"`])'
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

[JSON](https://github.com/k-kinzal/pr/blob/master/doc/spec.json)

See below for a detailed description of each item.

- [Pull Requests](https://developer.github.com/v3/pulls/)
- [Comments](https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request)
- [Reviews](https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request)
- [Commits](https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request)
- [Statuses](https://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref)

`"comments"`, `"reviews"`, `"commits"`, and `"statuses"` cannot be used by default with PR link relation.
If the parameter is required, specify option `--with-comments`, `--with-reviews`, `--with-commits`, `--with-statuses` or `--with-all`.

NOTE: `--with-all` options have very poor performance. Not recommended for uses other than debugging.

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

### GitHub Action

If you execute PR with [GitHub Action], the rules are automatically completed by event type.

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
head == `"[Branch Name]"`
```

Please see the [event trigger](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows) for details.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Fpr?ref=badge_large)