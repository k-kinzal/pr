package cmd

import (
	"os"
	"strings"

	"golang.org/x/xerrors"

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
				return xerrors.New("`owner/repo` argument is required")
			}
			s := strings.Split(args[0], "/")
			globalOption.Owner = s[0]
			globalOption.Repo = s[1]

			if globalOption.Token == "" {
				globalOption.Token = os.Getenv("GITHUB_TOKEN")
			}

			if globalOption.Token == "" {
				return xerrors.New("--token or GITHUB_TOKEN is required")
			}

			return nil
		},
		Version: GetVersion(),
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&globalOption.Token, "token", "", "personal access token to manipulate PR [GITHUB_TOKEN]")
	rootCmd.PersistentFlags().BoolVar(&exitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
}

func Execute() error {
	return rootCmd.Execute()
}
