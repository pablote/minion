package cmd

import (
	"github.com/pablote/minion/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update repos",
	Long:  `Update all develop and main/master branches to match origin in all repos`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn:    updateFn,
		}.Execute()
	},
}

func updateFn(path string) (string, error) {
	_, _, err := lib.RunCommand("git", path, "fetch", "--all", "--prune")
	if err != nil {
		return "", err
	}

	_, _, err = lib.RunCommand("git", path, "fetch", "--all", "--tags")
	if err != nil {
		return "", err
	}

	mainBranchName, err := lib.GetMainBranch(path)
	if err != nil {
		return "", err
	}

	branchesToUpdate := []string{mainBranchName, "develop"}
	for _, branch := range branchesToUpdate {
		hasBranch, err := lib.HasBranch(path, branch)
		if err != nil {
			return "", err
		}

		if hasBranch {
			_, _, err := lib.RunCommand("git", path, "checkout", branch)
			if err != nil {
				return "", err
			}

			_, _, err = lib.RunCommand("git", path, "pull", "origin", branch)
			if err != nil {
				return "", err
			}
		}
	}

	return "", nil
}
