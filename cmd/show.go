package cmd

import (
	"fmt"
	"os"

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
		opt.Rules = append(opt.Rules, fmt.Sprintf("number = `%d`", *pullNumber))
	}

	if branchName := action.BranchName(); branchName != nil {
		opt.Rules = append(opt.Rules, fmt.Sprintf("head = `\"%s\"`", *branchName))
	}

	// Do nothing
	//switch payload := action.Payload.(type) {
	//case ev.CheckRunPayload:
	//case ev.CheckSuitePayload:
	//case ev.CreatePayload:
	//case ev.DeletePayload:
	//case ev.DeploymentPayload:
	//case ev.DeploymentStatusPayload:
	//case ev.ForkPayload:
	//case ev.GollumPayload:
	//case ev.IssueCommentPayload:
	//case ev.IssuesPayload:
	//case ev.LabelPayload:
	//case ev.MemberPayload:
	//case ev.MilestonePayload:
	//case ev.PageBuildPayload:
	//case ev.ProjectPayload:
	//case ev.ProjectCardPayload:
	//case ev.ProjectColumnPayload:
	//case ev.PublicPayload:
	//case ev.PullRequestPayload:
	//case ev.PullRequestReviewPayload:
	//case ev.PullRequestReviewCommentPayload:
	//case ev.PushPayload:
	//case ev.ReleasePayload:
	//case ev.StatusPayload:
	//case ev.WatchPayload:
	//}

	if err := pr.Show(showOption); err != nil {
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
