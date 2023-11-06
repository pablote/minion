package lib

import (
	"strings"
)

func GetMainBranch(path string) (string, error) {
	response, _, err := RunCommand("git", path, "branch", "-a")
	if err != nil {
		return "", nil
	}

	for _, line := range strings.Split(response, "\n") {
		branchName := strings.TrimPrefix(strings.TrimSpace(line), "*")
		if branchName == "remotes/origin/master" || branchName == "master" {
			return "master", nil
		} else if branchName == "remotes/origin/main" || branchName == "main" {
			return "main", nil
		}
	}

	return "", nil
}
