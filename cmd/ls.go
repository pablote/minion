package cmd

import (
	"fmt"
	"github.com/pablote/minion/lib"
	"strings"

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
			Fn: func(path string) (string, error) {
				stdout, _, err := lib.RunCommand("ls", path, "-l")

				output := &strings.Builder{}
				_, _ = fmt.Fprintf(output, "%v:\n", path)
				_, _ = fmt.Fprintf(output, "%v\n", stdout)
				return output.String(), err
			},
		}.Execute()
	},
}
