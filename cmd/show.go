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
			if b, _ := cmd.Flags().GetBool("disable-all"); b {
				showOption.DisableComments = true
				showOption.DisableReviews = true
				showOption.DisableCommits = true
				showOption.DisableStatuses = true
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
	showCmd.Flags().BoolVar(&showOption.DisableComments, "disable-comments", false, "if true, do not retrieve comment link relations to PR")
	showCmd.Flags().BoolVar(&showOption.DisableReviews, "disable-reviews", false, "if true, do not retrieve review link relations to PR")
	showCmd.Flags().BoolVar(&showOption.DisableCommits, "disable-commits", false, "if true, do not retrieve commit link relations to PR")
	showCmd.Flags().BoolVar(&showOption.DisableStatuses, "disable-status", false, "if true, do not retrieve status link relations to PR")
	showCmd.Flags().Bool("disable-all", false, "if true, do not retrieve link relations to PR (NOTE: this option should be enabled if there are many PR)")
	showCmd.Flags().StringArrayVarP(&showOption.Rules, "rule", "l", nil, "JMESPath format view rules")
	rootCmd.AddCommand(showCmd)
}
