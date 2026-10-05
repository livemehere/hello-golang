package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	value := <-ch

	fmt.Printf("value : %d", value)
}
