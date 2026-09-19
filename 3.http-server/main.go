package main

import "fmt"

const (
	A = iota
	B
	C
)

func sum[T int | float32](nums ...T) T {
	var total T = 0
	for _, n := range nums {
		total += n
	}
	return total
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println(A, B, C)
}
