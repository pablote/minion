package lib

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

const concurrencyLevel = 4

type Runner struct {
	Paths []string
	Fn    func(path string) (string, error)
}

func (r Runner) Execute() {
	outputCh := make(chan string, len(r.Paths))
	wg := &sync.WaitGroup{}
	wg.Add(len(r.Paths))
	s := semaphore.NewWeighted(concurrencyLevel)
	c := context.Background()

	// run fn for all paths
	for _, path := range r.Paths {
		go func(path string) {
			err := s.Acquire(c, 1)
			if err != nil {
				panic(err)
			}

			output, err := r.Fn(path)

			if err != nil {
				fmt.Printf("Failed for %v\n", path)
				fmt.Println(err)
			} else {
				outputCh <- output
				fmt.Printf("Done for %v\n", path)
			}

			wg.Done()
			s.Release(1)
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
