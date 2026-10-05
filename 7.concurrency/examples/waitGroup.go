package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("worker", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	const workerCount = 5

	var wg sync.WaitGroup

	wg.Add(workerCount)
	for i := range workerCount {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("all done!")
}
