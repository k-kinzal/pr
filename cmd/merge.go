package cmd

import (
	"fmt"
	"os"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			mergeOption.Option = globalOption
			if err := pr.Merge(mergeOption); err != nil {
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
	}
)

func init() {
	mergeCmd.Flags().IntVar(&mergeOption.Limit, "limit", 100, "limit the number of merge")
	mergeCmd.Flags().IntVar(&mergeOption.Rate, "rate", 10, "API call seconds rate limit")
	mergeCmd.Flags().StringArrayVarP(&mergeOption.Rules, "rule", "l", nil, "JMESPath format merge rules")
	mergeCmd.Flags().StringVar(&mergeOption.CommitTitleTemplate, "commit-title", "Merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}", "title for the automatic commit message.")
	mergeCmd.Flags().StringVar(&mergeOption.CommitMessageTemplate, "commit-message", "{{ .Title }}", "extra detail to append to automatic commit message")
	mergeCmd.Flags().StringVar(&mergeOption.MergeMethod, "method", "merge", "merge method to use. possible values are merge, squash or rebase")
	rootCmd.AddCommand(mergeCmd)
}
