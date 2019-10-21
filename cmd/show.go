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
	showCmd.Flags().StringArrayVarP(&showOption.Rules, "rule", "l", nil, "JMESPath format view rules")
	rootCmd.AddCommand(showCmd)
}
