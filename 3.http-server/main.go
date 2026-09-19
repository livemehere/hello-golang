package main

import "fmt"

func sum(ch chan int) {
	fmt.Println("worker")
	ch <- 10
}

func main() {
	ch := make(chan int)

	fmt.Println("before")

	go sum(ch)

	v := <-ch

	fmt.Println("after", v)
}
