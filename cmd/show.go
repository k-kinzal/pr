package cmd

import (
	"fmt"
	"os"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			showOption.Option = globalOption
			if b, _ := cmd.Flags().GetBool("with-all"); b {
				showOption.EnableComments = true
				showOption.EnableReviews = true
				showOption.EnableCommits = true
				showOption.EnableStatuses = true
			}
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
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

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
