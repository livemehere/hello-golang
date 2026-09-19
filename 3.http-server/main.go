package main

import (
	"fmt"
	"sync"
	"time"
)

func say(word string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		fmt.Printf("%d : %s\n", i, word)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go say("hello", &wg)
	go say("world", &wg)

	wg.Wait()
}
