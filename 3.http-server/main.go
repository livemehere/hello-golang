package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3

		close(ch)
	}()

	read := func() {
		v, ok := <-ch
		fmt.Println(v, ok)
	}

	read()
	read()
	read()
	read()
}
