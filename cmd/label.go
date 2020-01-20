package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func LabelRun(cmd *cobra.Command, args []string) error {
	labelOption.Action = pr.LabelAction(labelAction)

	pulls, err := pr.Label(owner, repo, labelOption)
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
	labelOption *pr.LabelOption
	labelCmd    = &cobra.Command{
		Use:           "label owner/repo",
		Short:         "Manipulate labels that match a rule",
		RunE:          LabelRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	labelAction string
)

func setLabelFrags(cmd *cobra.Command) *pr.LabelOption {
	opt := &pr.LabelOption{}
	cmd.Flags().StringArrayVarP(&opt.Labels, "label", "", nil, "Manipulate labels")
	cmd.Flags().StringVar(&labelAction, "action", "append", "Manipulation of `append`, `remove`, `replace` for labels")
	return opt
}

func init() {
	labelOption = setLabelFrags(labelCmd)
	labelOption.ListOption = setListFrags(labelCmd)
	rootCmd.AddCommand(labelCmd)
}
