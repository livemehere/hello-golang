package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count atomic.Int64

	wg.Add(1000)

	for range 1000 {
		go func() {
			defer wg.Done()
			count.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(count.Load())
}
