package cmd

import (
	"fmt"
	"minion/lib"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(releasesCmd)
}

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Show open releases",
	Long:  `Show open releases`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn:    releasesFn,
		}.Execute()
	},
}

func releasesFn(path string) (string, error) {
	argsList := [][]string{
		{"branch"},
		{"branch", "-r"},
	}

	branches := make([]string, 0)

	for _, args := range argsList {
		response, _, err := lib.RunCommand("git", path, args...)
		if err != nil {
			return "", err
		}

		for _, branch := range strings.Split(response, "\n") {
			if strings.Index(branch, "release/") > -1 {
				branches = append(branches, strings.TrimSpace(branch))
			}
		}
	}

	output := &strings.Builder{}
	if len(branches) > 0 {
		_, _ = fmt.Fprintf(output, "%v:", path)
		for _, branch := range branches {
			_, _ = fmt.Fprintf(output, "\t%v:\n", branch)
		}
	}

	return output.String(), nil
}
