package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/k-kinzal/pr/pkg/pr"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func CheckRun(cmd *cobra.Command, args []string) error {
	if b, _ := cmd.Flags().GetBool("merge"); b {
		checkOption.Action = "merge"
	}
	pulls, err := pr.Check(owner, repo, checkOption)
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
	checkOption *pr.CheckOption
	checkCmd    = &cobra.Command{
		Use:   "check owner/repo",
		Short: "Check if PR matches the rule and change PR status",
		Long: `Check if PR matches the rule and change PR status
Be sure to specify 'number', 'head.ref', or 'head.shas' rules with this command. Check if the target PR matches all the rules
`,
		RunE:          CheckRun,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func setCheckFrags(cmd *cobra.Command) *pr.CheckOption {
	opt := &pr.CheckOption{}
	cmd.Flags().StringVar(&opt.TargetURL, "target-url", "", "URL linked from the check result displayed on GitHub. If you are running with CI, specify that URL (default \"https://github.com/{{ .Owner }}/{{ .Repo }}/commit/{{ .Head.Sha }}/checks\" in GitHub Action)")
	cmd.Flags().Bool("merge", false, "If the check is successful, merge the PR")
	return opt
}

func init() {
	checkOption = setCheckFrags(checkCmd)
	checkOption.MergeOption = setMergeFrags(checkCmd)
	checkOption.ListOption = setListFrags(checkCmd)
	rootCmd.AddCommand(checkCmd)
}
