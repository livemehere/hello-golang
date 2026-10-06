package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	done := make(chan bool)

	var counter Counter

	created := 0

loop:
	for range 1000 {
		select {
		case <-done:
			break loop
		default:
		}

		created++

		wg.Add(1)
		go func() {
			defer wg.Done()

			counter.mu.Lock()

			if counter.value == 500 {
				counter.mu.Unlock()
				once.Do(func() {
					close(done)
				})
				return
			}
			counter.value++
			counter.mu.Unlock()
		}()

	}

	wg.Wait()

	fmt.Println(counter.value, created)
}
