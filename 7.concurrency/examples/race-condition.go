package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	count := 0

	wg.Add(1000)

	for range 1000 {
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
