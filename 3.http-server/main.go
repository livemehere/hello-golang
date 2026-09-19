package main

import (
	"fmt"
	"time"
)

type empty struct{}

func main() {
	done := make(chan empty)

	say := func(word string) {
		for i := range 10 {
			fmt.Printf("%d : %s\n", i, word)
			time.Sleep(time.Millisecond * 100)
		}
		done <- empty{}
	}

	go say("hello")
	go say("world")

	<-done
	<-done
}
