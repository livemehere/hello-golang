package main

import (
	"fmt"

	"my-server/internal/math"
)

func main() {
	var n int = 10
	var s string = "Hello world"
	var ok bool = false

	fmt.Println(n, s, ok)

	fmt.Println(math.Add(1, 2))
}
