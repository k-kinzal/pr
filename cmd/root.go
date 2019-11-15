package cmd

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/xerrors"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

var (
	globalOption pr.Option
	exitCode     bool
	noExitCode   bool
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
				return xerrors.New("GITHUB_TOKEN or --token arguments is required")
			}

			return nil
		},
		Version:       GetVersion(),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&globalOption.Token, "token", "", "personal access token to manipulate PR [GITHUB_TOKEN]")
	rootCmd.PersistentFlags().BoolVar(&exitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	rootCmd.PersistentFlags().BoolVar(&noExitCode, "no-exit-code", false, "always returns 0 even if an error occurs")
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("pr: %s", err))
		if noExitCode {
			os.Exit(0)
		}
		switch err.(type) {
		case *pr.NoMatchError:
			if exitCode {
				os.Exit(127)
			}
		}
		os.Exit(1)
	}
	return nil
}
