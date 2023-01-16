package lib

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func RunCommand(name, dir string, args ...string) (string, string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return stdout.String(), stderr.String(), errors.New(fmt.Sprintf("%s: %s (failed on command: %s)", err.Error(), stderr, args))
	}

	return stdout.String(), stderr.String(), nil
}
