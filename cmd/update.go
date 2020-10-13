package cmd

import (
	"fmt"
	"run-cli/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all repos",
	Long:  `Update all develop and master branches to match origin in all repos`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn:    updateFn,
		}.Execute()
	},
}

func updateFn(path string) error {
	_, _, err := lib.RunCommand("git", path, "fetch", "--all", "--prune")
	if err != nil {
		return err
	}

	branchesToUpdate := []string{"master", "develop"}
	for _, branch := range branchesToUpdate {
		hasBranch, err := hasBranch(path, branch)
		if err != nil {
			return err
		}

		if hasBranch {
			_, _, err := lib.RunCommand("git", path, "checkout", branch)
			if err != nil {
				return err
			}

			_, _, err = lib.RunCommand("git", path, "pull", "origin", branch)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func hasBranch(path, branch string) (bool, error) {
	response, _, err := lib.RunCommand("git", path, "show-ref", fmt.Sprintf("refs/heads/%v", branch))
	if err != nil {
		return false, nil
	}

	if len(response) > 0 {
		return true, nil
	}

	return false, nil
}
