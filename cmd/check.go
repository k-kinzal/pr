package cmd

import (
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"

	"github.com/k-kinzal/pr/pkg/action"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

func CheckRun(cmd *cobra.Command, args []string) error {
	opt := checkOption
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

	if b, _ := cmd.Flags().GetBool("merge"); b {
		opt.Action = "merge"
	}

	if err := pr.Check(opt); err != nil {
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

var (
	checkOption pr.CheckOption
	checkCmd    = &cobra.Command{
		Use:   "check owner/repo",
		Short: "Check if PR matches the rule and change PR status",
		Long: `Check if PR matches the rule and change PR status
Be sure to specify 'number', 'head.ref', or 'head.shas' rules with this command. Check if the target PR matches all the rules
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
			}
			return nil
		},
		RunE:          CheckRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	checkCmd.Flags().IntVar(&checkOption.Limit, "limit", 100, "limit the number of views")
	checkCmd.Flags().IntVar(&checkOption.Rate, "rate", 10, "API call seconds rate limit")
	checkCmd.Flags().BoolVar(&checkOption.EnableComments, "with-comments", false, "if true, do retrieve comment link relations to PR")
	checkCmd.Flags().BoolVar(&checkOption.EnableReviews, "with-reviews", false, "if true, do retrieve review link relations to PR")
	checkCmd.Flags().BoolVar(&checkOption.EnableCommits, "with-commits", false, "if true, do retrieve commit link relations to PR")
	checkCmd.Flags().BoolVar(&checkOption.EnableStatuses, "with-statuses", false, "if true, do retrieve status link relations to PR")
	checkCmd.Flags().Bool("with-all", false, "if true, do retrieve link relations to PR (NOTE: this option should be disabled if there are many PR)")
	checkCmd.Flags().StringArrayVarP(&checkOption.Rules, "rule", "l", nil, "JMESPath format view rules")
	checkCmd.Flags().StringVar(&checkOption.TargetURL, "target-url", "", "URL linked from the check result displayed on GitHub. If you are running with CI, specify that URL (default \"https://github.com/{{ .Owner }}/{{ .Repo }}/commit/{{ .Head.Sha }}/checks\" in GitHub Action)")
	checkCmd.Flags().Bool("merge", false, "If the check is successful, merge the PR")
	checkCmd.Flags().StringVar(&checkOption.MergeOption.CommitTitleTemplate, "commit-title", "merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}", "title for the automatic commit message.")
	checkCmd.Flags().StringVar(&checkOption.MergeOption.CommitMessageTemplate, "commit-message", "{{ .Title }}", "extra detail to append to automatic commit message")
	checkCmd.Flags().StringVar(&checkOption.MergeOption.MergeMethod, "method", "merge", "merge method to use. possible values are merge, squash or rebase")
	rootCmd.AddCommand(checkCmd)
}
