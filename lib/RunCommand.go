package lib

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func RunCommand(name, dir string, args ...string) (string, string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	//log.Println("[INFO] Run command: ", name, args)

	if err := cmd.Run(); err != nil {
		log.Println(fmt.Sprintf("[DEBUG] Run command: stdout: %v", stdout.String()))
		log.Println(fmt.Sprintf("[DEBUG] Run command: stderr: %v", stderr.String()))
		return stdout.String(), stderr.String(), err
	}

	//log.Println(fmt.Sprintf("[DEBUG] Run command: stdout: %v", stdout.String()))
	//log.Println(fmt.Sprintf("[DEBUG] Run command: stderr: %v", stderr.String()))

	return stdout.String(), stderr.String(), nil
}
