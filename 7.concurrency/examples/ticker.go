package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "done!"
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println("value :", v)

		case <-ticker.C:
			fmt.Println("tick...")
		}
	}
}
