package cmd

import (
	"fmt"
	"github.com/pablote/minion/lib"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(latestCmd)
}

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Get latest tag",
	Long:  `Get the latest tagged version for each repo`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		paths := lib.GetPaths(args)
		lib.Runner{
			Paths: paths,
			Fn:    latestFn,
		}.Execute()
	},
}

func latestFn(path string) (string, error) {
	response, _, err := lib.RunCommand("git", path, "tag")
	if err != nil {
		return "", err
	}

	tags := strings.Split(response, "\n")

	versionTags := make([]string, 0)
	for _, tag := range tags {
		if strings.HasPrefix(tag, "v") {
			versionTags = append(versionTags, tag)
		}
	}

	latest := ""
	if len(versionTags) >= 1 {
		latest = versionTags[len(versionTags)-1]
	}

	return fmt.Sprintf("%s: %s\n", path, latest), nil
}
