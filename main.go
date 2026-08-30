package main

import (
	"fmt"
)

// types of args
func add(x, y int) int {
	return x + y
}

// multiple return
func swap(a, b string) (string, string) {
	return b, a
}

// naked return
func split(v int) (x, y int) {
	x = v / 3
	y = v % 3
	return
}

func main() {
	fmt.Println(split(10))
}
