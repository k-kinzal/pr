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
	token      string
	owner      string
	repo       string
	exitCode   bool
	noExitCode bool
	rate       int
	rootCmd    = &cobra.Command{
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
			owner = s[0]
			repo = s[1]

			if token == "" {
				token = os.Getenv("GITHUB_TOKEN")
				if token == "" {
					return xerrors.New("GITHUB_TOKEN or --token arguments is required")
				}
			}
			pr.SetToken(token)

			if rate, _ := cmd.Flags().GetInt("rate"); rate < 1 {
				return xerrors.New("specify 1 or more for --rate.")
			}

			return nil
		},
		Version:       GetVersion(),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "personal access token to manipulate PR [GITHUB_TOKEN]")
	rootCmd.PersistentFlags().BoolVar(&exitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	rootCmd.PersistentFlags().BoolVar(&noExitCode, "no-exit-code", false, "always returns 0 even if an error occurs")
	rootCmd.PersistentFlags().IntVar(&rate, "rate", 10, "API call seconds rate limit")
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if noExitCode {
			os.Exit(0)
		}
		os.Exit(1)
	}
	return nil
}
