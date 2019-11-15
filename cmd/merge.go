package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func MergeRun(cmd *cobra.Command, args []string) error {
	pulls, err := pr.Merge(owner, repo, mergeOption)
	if err != nil {
		if _, ok := err.(*pr.NoMatchError); ok {
			fmt.Fprintln(os.Stderr, err.Error())
			fmt.Fprintln(os.Stdout, "[]")
			if exitCode {
				os.Exit(127)
			}
			return nil
		}
		return err
	}

	out, err := json.Marshal(pulls)
	if err != nil {
		return xerrors.Errorf("merge: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}

var (
	mergeOption *pr.MergeOption
	mergeCmd    = &cobra.Command{
		Use:           "merge owner/repo",
		Short:         "Merge PR that matches a rule",
		RunE:          MergeRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func setMergeFrags(cmd *cobra.Command) *pr.MergeOption {
	opt := &pr.MergeOption{}
	cmd.Flags().StringVar(&opt.CommitTitleTemplate, "commit-title", "Merge pull request #{{ .Number }} from {{ .Owner }}/{{ .Head.Ref }}", "title for the automatic commit message.")
	cmd.Flags().StringVar(&opt.CommitMessageTemplate, "commit-message", "{{ .Title }}", "extra detail to append to automatic commit message")
	cmd.Flags().StringVar(&opt.MergeMethod, "method", "merge", "merge method to use. possible values are merge, squash or rebase")
	return opt
}

func init() {
	mergeOption = setMergeFrags(mergeCmd)
	mergeOption.ListOption = setListFrags(mergeCmd)
	rootCmd.AddCommand(mergeCmd)
}
