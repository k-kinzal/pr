package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	githubToken string
	rootCmd     = &cobra.Command{
		Use:   "pr",
		Short: "PR operates multiple Pull Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&githubToken, "token", os.Getenv("GITHUB_TOKEN"), "personal access token to manipulate PR [GITHUB_TOKEN]")
}

func Execute() error {
	return rootCmd.Execute()
}
