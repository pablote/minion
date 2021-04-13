package lib

import (
	"fmt"
	"strings"
)

type Runner struct {
	Paths []string
	Fn    func(path string) (string, error)
}

func (r Runner) Execute() {
	accumulatedOutput := &strings.Builder{}

	for _, path := range r.Paths {
		output, err := r.Fn(path)

		if err != nil {
			fmt.Printf("Failed for %v\n", path)
			fmt.Println(err)
		} else {
			_, _ = fmt.Fprint(accumulatedOutput, output)
			fmt.Printf("Done for %v\n", path)
		}
	}

	fmt.Printf("\n%v\n", accumulatedOutput.String())
}
