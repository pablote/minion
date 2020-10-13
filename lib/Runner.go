package lib

import "fmt"

type Runner struct {
	Paths []string
	Fn func(path string) error
}

func (r Runner) Execute() {
	for _, path := range r.Paths {
		err := r.Fn(path)

		if err != nil {
			fmt.Printf("Failed for %v\n", path)
			fmt.Println(err)
		} else {
			fmt.Printf("Done for %v\n", path)
		}
	}
}
