package cmd

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/pablote/minion/lib"
	"sort"
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

	// convert response into a list of strings
	tags := strings.Split(response, "\n")

	// strip v prefix
	versionTags := make([]string, 0)
	for _, tag := range tags {
		if strings.HasPrefix(tag, "v") {
			versionTags = append(versionTags, tag)
		}
	}

	// sort based on sem ver
	sort.Slice(versionTags, func(i, j int) bool {
		v1, err := version.NewVersion(versionTags[i])
		if err != nil {
			fmt.Println("failed to parse version %s", versionTags[i])
			return false
		}

		v2, err := version.NewVersion(versionTags[j])
		if err != nil {
			fmt.Println("failed to parse version %s", versionTags[j])
			return false
		}

		return v1.LessThan(v2)
	})

	// get the latest
	latest := ""
	if len(versionTags) >= 1 {
		latest = versionTags[len(versionTags)-1]
	}

	return fmt.Sprintf("%s: %s\n", path, latest), nil
}
