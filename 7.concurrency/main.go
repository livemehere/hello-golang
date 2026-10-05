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
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
	fmt.Println("all done!")
}
