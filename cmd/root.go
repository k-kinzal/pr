package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

var (
	globalOption pr.Option
	exitCode     bool
	rootCmd      = &cobra.Command{
		Use:   "pr",
		Short: "PR operates multiple Pull Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("`owner/repo` argument is required")
			}
			s := strings.Split(args[0], "/")
			globalOption.Owner = s[0]
			globalOption.Repo = s[1]
			return nil
		},
		Version: GetVersion(),
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&globalOption.Token, "token", os.Getenv("GITHUB_TOKEN"), "personal access token to manipulate PR [GITHUB_TOKEN]")
	rootCmd.PersistentFlags().BoolVar(&exitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
}

func Execute() error {
	return rootCmd.Execute()
}
