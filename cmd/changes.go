package cmd

import (
	"fmt"
	"github.com/pablote/minion/lib"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changesCmd)
}

var changesCmd = &cobra.Command{
	Use:   "changes",
	Short: "Show changes between develop and main/master",
	Long:  `Show changes between develop and main/master`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn:    changesFn,
		}.Execute()
	},
}

func changesFn(path string) (string, error) {
	mainBranchName, err := lib.GetMainBranch(path)
	if err != nil {
		return "", err
	}

	hasMaster, err := lib.HasBranch(path, mainBranchName)
	if err != nil {
		return "", err
	}

	hasDevelop, err := lib.HasBranch(path, "develop")
	if err != nil {
		return "", err
	}

	output := &strings.Builder{}
	if hasMaster && hasDevelop {
		response, _, err := lib.RunCommand("git", path, "log", "--pretty=oneline", "--no-merges", "develop", fmt.Sprintf("^%s", mainBranchName))
		if err != nil {
			return "", err
		}

		if len(response) > 0 {
			_, _ = fmt.Fprintf(output, "Changes for %v:\n", path)
			_, _ = fmt.Fprintf(output, "%v\n", response)
		}
	}

	return output.String(), nil
}
