package cmd

import (
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"

	"github.com/k-kinzal/pr/pkg/action"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

var (
	mergeOption pr.MergeOption
	mergeCmd    = &cobra.Command{
		Use:   "merge owner/repo",
		Short: "Merge PR that matches a rule",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
			}
			return nil
		},
		RunE:          MergeRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func MergeRun(cmd *cobra.Command, args []string) error {
	opt := mergeOption
	opt.Option = globalOption
	if b, _ := cmd.Flags().GetBool("with-all"); b {
		opt.EnableComments = true
		opt.EnableReviews = true
		opt.EnableCommits = true
		opt.EnableStatuses = true
	}

	if pullNumber := action.PullNumber(); pullNumber != nil {
		opt.Rules = append(opt.Rules, fmt.Sprintf("number == `%d`", *pullNumber))
	}

	if branchName := action.BranchName(); branchName != nil {
		opt.Rules = append(opt.Rules, fmt.Sprintf("head.ref == `\"%s\"`", *branchName))
	}

	switch action.Payload.(type) {
	case github.PageBuildEvent:
		opt.Rules = append(opt.Rules, fmt.Sprintf("head.sha == `\"%s\"`", action.SHA))
	case github.StatusEvent:
		opt.Rules = append(opt.Rules, fmt.Sprintf("head.sha == `\"%s\"`", action.SHA))
	}

	if err := pr.Merge(opt); err != nil {
		switch err.(type) {
		case *pr.NoMatchError:
			if exitCode {
				os.Exit(127)
			}
			return nil
		}
		return err
	}

	return nil
}

func init() {
	mergeCmd.Flags().IntVar(&mergeOption.Limit, "limit", 100, "limit the number of merge")
	mergeCmd.Flags().IntVar(&mergeOption.Rate, "rate", 10, "API call seconds rate limit")
	mergeCmd.Flags().BoolVar(&mergeOption.EnableComments, "with-comments", false, "if true, do retrieve comment link relations to PR")
	mergeCmd.Flags().BoolVar(&mergeOption.EnableReviews, "with-reviews", false, "if true, do retrieve review link relations to PR")
	mergeCmd.Flags().BoolVar(&mergeOption.EnableCommits, "with-commits", false, "if true, do retrieve commit link relations to PR")
	mergeCmd.Flags().BoolVar(&mergeOption.EnableStatuses, "with-statuses", false, "if true, do retrieve status link relations to PR")
	mergeCmd.Flags().Bool("with-all", false, "if true, do retrieve link relations to PR (NOTE: this option should be disabled if there are many PR)")
	mergeCmd.Flags().StringArrayVarP(&mergeOption.Rules, "rule", "l", nil, "JMESPath format merge rules")
	mergeCmd.Flags().StringVar(&mergeOption.CommitTitleTemplate, "commit-title", "Merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}", "title for the automatic commit message.")
	mergeCmd.Flags().StringVar(&mergeOption.CommitMessageTemplate, "commit-message", "{{ .Title }}", "extra detail to append to automatic commit message")
	mergeCmd.Flags().StringVar(&mergeOption.MergeMethod, "method", "merge", "merge method to use. possible values are merge, squash or rebase")
	rootCmd.AddCommand(mergeCmd)
}
