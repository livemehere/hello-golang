package main

import "fmt"

func worker(ch chan int) {
	fmt.Println("worker: before send")

	ch <- 100

	fmt.Println("worker: after send")
}

func main() {
	ch := make(chan int)

	go worker(ch)

	fmt.Println("main: before receive")

	value := <-ch

	fmt.Println("main: received", value)
}
