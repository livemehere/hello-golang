package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "fast"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "slow"
	}()

	for range 2 {
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)

		case v := <-ch2:
			fmt.Println("ch2:", v)
		}
	}
}
