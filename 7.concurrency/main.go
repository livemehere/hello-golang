package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var state atomic.Int64

	state.Store(0)

	ok := state.CompareAndSwap(0, 99)

	fmt.Println(ok)
	fmt.Println(state.Load())
}
