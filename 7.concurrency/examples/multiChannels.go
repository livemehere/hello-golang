package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- id * id
}

func main() {
	ch := make(chan int, 2)

	var wg sync.WaitGroup

	const workerCount = 5

	wg.Add(workerCount)
	for i := range workerCount {
		go worker(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
