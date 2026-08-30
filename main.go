package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
