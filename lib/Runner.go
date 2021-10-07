package lib

import (
	"fmt"
	"sync"
)

type Runner struct {
	Paths []string
	Fn    func(path string) (string, error)
}

func (r Runner) Execute() {
	outputCh := make(chan string, len(r.Paths))

	wg := &sync.WaitGroup{}
	wg.Add(len(r.Paths))

	// run fn for all paths
	for _, path := range r.Paths {
		go func(path string) {
			fmt.Printf("Starting for %v\n", path)
			output, err := r.Fn(path)

			if err != nil {
				fmt.Printf("Failed for %v\n", path)
				fmt.Println(err)
			} else {
				outputCh <- output
				fmt.Printf("Done for %v\n", path)
			}

			wg.Done()
		}(path)
	}

	// wait for tasks
	wg.Wait()
	close(outputCh)

	// print output
	for output := range outputCh {
		fmt.Printf("%s", output)
	}
}
