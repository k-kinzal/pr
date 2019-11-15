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
	showOption pr.PullOption
	showCmd    = &cobra.Command{
		Use:   "show owner/repo",
		Short: "Show PR that matches a rule",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
			}
			return nil
		},
		RunE:          ShowRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func ShowRun(cmd *cobra.Command, args []string) error {
	opt := showOption
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

	if err := pr.Show(opt); err != nil {
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
	showCmd.Flags().IntVar(&showOption.Limit, "limit", 100, "limit the number of views")
	showCmd.Flags().IntVar(&showOption.Rate, "rate", 10, "API call seconds rate limit")
	showCmd.Flags().BoolVar(&showOption.EnableComments, "with-comments", false, "if true, do retrieve comment link relations to PR")
	showCmd.Flags().BoolVar(&showOption.EnableReviews, "with-reviews", false, "if true, do retrieve review link relations to PR")
	showCmd.Flags().BoolVar(&showOption.EnableCommits, "with-commits", false, "if true, do retrieve commit link relations to PR")
	showCmd.Flags().BoolVar(&showOption.EnableStatuses, "with-statuses", false, "if true, do retrieve status link relations to PR")
	showCmd.Flags().Bool("with-all", false, "if true, do retrieve link relations to PR (NOTE: this option should be disabled if there are many PR)")
	showCmd.Flags().StringArrayVarP(&showOption.Rules, "rule", "l", nil, "JMESPath format view rules")
	rootCmd.AddCommand(showCmd)
}
