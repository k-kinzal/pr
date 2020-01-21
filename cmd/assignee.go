package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func AssigneeRun(cmd *cobra.Command, args []string) error {
	assigneeOption.Action = pr.AssigneeAction(assigneeAction)

	pulls, err := pr.Assignee(owner, repo, assigneeOption)
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
	assigneeOption *pr.AssigneeOption
	assigneeCmd    = &cobra.Command{
		Use:           "assignee owner/repo",
		Short:         "Manipulate assignees that match a rule",
		RunE:          AssigneeRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	assigneeAction string
)

func setAssigneeFrags(cmd *cobra.Command) *pr.AssigneeOption {
	opt := &pr.AssigneeOption{}
	cmd.Flags().StringArrayVarP(&opt.Assignees, "assignee", "", nil, "Manipulate assignee")
	cmd.Flags().StringVar(&assigneeAction, "action", "append", "Manipulation of `append`, `remove`, `replace` for assignees")
	return opt
}

func init() {
	assigneeOption = setAssigneeFrags(assigneeCmd)
	assigneeOption.ListOption = setListFrags(assigneeCmd)
	rootCmd.AddCommand(assigneeCmd)
}
