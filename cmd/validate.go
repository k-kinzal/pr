package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/api"

	"golang.org/x/xerrors"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
)

func ValidateRun(cmd *cobra.Command, args []string) error {
	result, pulls := pr.Validate(owner, repo, validateOption)
	if pulls == nil {
		pulls = make([]*api.PullRequest, 0)
	}
	for _, res := range result {
		if res.Success() {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("[X] %s", res.String()))
		} else {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("[ ] %s", res.String()))
		}
	}

	out, err := json.Marshal(pulls)
	if err != nil {
		return xerrors.Errorf("show: %s", err)
	}
	fmt.Fprintln(os.Stdout, string(out))

	return nil
}

var (
	validateOption *pr.ListOption
	validateCmd    = &cobra.Command{
		Use:           "validate owner/repo",
		Short:         "Validate the rules",
		RunE:          ValidateRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	validateOption = setListFrags(validateCmd)
	rootCmd.AddCommand(validateCmd)
}
