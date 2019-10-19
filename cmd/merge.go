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
	mergeLimit                 int
	mergeRules                 []string
	mergeExitCode              bool
	mergeCommitTitleTemplate   string
	mergeCommitMessageTemplate string
	mergeMethod                string
	mergeCmd                   = &cobra.Command{
		Use:   "merge owner/repo",
		Short: "Merge PR that matches a rule",
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
			opt := &pr.MergeOption{
				CommitTitleTemplate:   mergeCommitTitleTemplate,
				CommitMessageTemplate: mergeCommitMessageTemplate,
				MergeMethod:           mergeMethod,
				PullOption: &pr.PullOption{
					Limit: mergeLimit,
					Rules: mergeRules,
					Option: &pr.Option{
						Token: githubToken,
						Owner: s[0],
						Repo:  s[1],
					},
				},
			}
			if err := pr.Merge(opt); err != nil {
				switch err.(type) {
				case *pr.NoMatchError:
					if mergeExitCode {
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
	mergeCmd.Flags().IntVar(&mergeLimit, "limit", 100, "limit the number of merge")
	mergeCmd.Flags().StringArrayVarP(&mergeRules, "rule", "l", nil, "JMESPath format merge rules")
	mergeCmd.Flags().BoolVar(&mergeExitCode, "exit-code", false, "returns an exit code of 127 if no PR matches the rule")
	mergeCmd.Flags().StringVar(&mergeCommitTitleTemplate, "commit-title", "Merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}", "title for the automatic commit message.")
	mergeCmd.Flags().StringVar(&mergeCommitMessageTemplate, "commit-message", "{{ .Title }}", "extra detail to append to automatic commit message")
	mergeCmd.Flags().StringVar(&mergeMethod, "method", "merge", "merge method to use. possible values are merge, squash or rebase")
	rootCmd.AddCommand(mergeCmd)
}
