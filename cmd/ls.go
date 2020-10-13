package cmd

import (
	"run-cli/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files",
	Long:  `List files`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn: func(path string) error {
				_, _, err := lib.RunCommand("ls", path, "-l")
				return err
			},
		}.Execute()
	},
}
