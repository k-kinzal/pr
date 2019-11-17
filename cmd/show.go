package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/xerrors"

	"github.com/google/go-github/v28/github"
	"github.com/k-kinzal/pr/pkg/action"
	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

func ShowRun(cmd *cobra.Command, args []string) error {
	pulls, err := pr.List(owner, repo, listOption)
	if err != nil {
		if _, ok := err.(*pr.NoMatchError); ok {
			fmt.Fprintln(os.Stderr, err.Error())
			fmt.Fprintln(os.Stdout, "[]")
			if exitCode {
				os.Exit(127)
			}
			return nil
		}
		return err
	}

	out, err := json.Marshal(pulls)
	if err != nil {
		return xerrors.Errorf("show: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}

var (
	listOption *pr.ListOption
	showCmd    = &cobra.Command{
		Use:           "show owner/repo",
		Short:         "Show PR that matches a rule",
		RunE:          ShowRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func setListFrags(cmd *cobra.Command) *pr.ListOption {
	opt := &pr.ListOption{}
	cmd.Flags().IntVar(&opt.Limit, "limit", 100, "limit the number of views")
	cmd.Flags().BoolVar(&opt.EnableComments, "with-comments", false, "if true, do retrieve comment link relations to PR")
	cmd.Flags().BoolVar(&opt.EnableReviews, "with-reviews", false, "if true, do retrieve review link relations to PR")
	cmd.Flags().BoolVar(&opt.EnableCommits, "with-commits", false, "if true, do retrieve commit link relations to PR")
	cmd.Flags().BoolVar(&opt.EnableStatuses, "with-statuses", false, "if true, do retrieve status link relations to PR")
	cmd.Flags().BoolVar(&opt.EnableChecks, "with-checks", false, "if true, do retrieve check link relations to PR")
	cmd.Flags().Bool("with-all", false, "if true, do retrieve link relations to PR (NOTE: this option should be disabled if there are many PR)")
	cmd.Flags().StringArrayVarP(&opt.Rules, "rule", "l", nil, "JMESPath format view rules")
	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		opt.Rate = rate

		if b, _ := cmd.Flags().GetBool("with-all"); b {
			opt.EnableComments = true
			opt.EnableReviews = true
			opt.EnableCommits = true
			opt.EnableStatuses = true
			opt.EnableChecks = true
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
	}
	return opt
}

func init() {
	listOption = setListFrags(showCmd)
	rootCmd.AddCommand(showCmd)
}
