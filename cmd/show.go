package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

var (
	showLimit    int
	showRules    []string
	showExitCode bool
	showCmd      = &cobra.Command{
		Use:   "show owner/repo [pr number]",
		Short: "Show PR that matches a rule",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
			}
			r, _ := regexp.Compile(`^[^/]+/[^/]+$`)
			if !r.MatchString(args[0]) {
				return fmt.Errorf("invalid arguments \"%s\"", args[0])
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s := strings.Split(args[0], "/")
			opt := &pr.PullOption{
				Limit: showLimit,
				Rules: showRules,
				Option: &pr.Option{
					Token: githubToken,
					Owner: s[0],
					Repo:  s[1],
				},
			}
			if err := pr.Show(opt); err != nil {
				switch err.(type) {
				case *pr.NoMatchError:
					if showExitCode {
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
	showCmd.Flags().IntVar(&showLimit, "limit", 100, "limit the number of views")
	showCmd.Flags().StringArrayVarP(&showRules, "rule", "l", nil, "JMESPath format view rules")
	showCmd.Flags().BoolVar(&showExitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	rootCmd.AddCommand(showCmd)
}
