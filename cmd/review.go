package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func ReviewRun(cmd *cobra.Command, args []string) error {
	pulls, err := pr.Review(owner, repo, reviewOption)
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
		return xerrors.Errorf("review: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}

var (
	reviewOption *pr.ReviewOption
	reviewCmd    = &cobra.Command{
		Use:           "review owner/repo",
		Short:         "Add review to PRs that match rules",
		RunE:          ReviewRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func setReviewFrags(cmd *cobra.Command) *pr.ReviewOption {
	opt := &pr.ReviewOption{}
	cmd.Flags().StringVar(&opt.Action, "action", "approve", "review action to take. currently, approve is only supported")
	return opt
}

func init() {
	reviewOption = setReviewFrags(reviewCmd)
	reviewOption.ListOption = setListFrags(reviewCmd)
	rootCmd.AddCommand(reviewCmd)
}
