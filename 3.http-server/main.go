package main

import (
	"fmt"
)

func sum[T int | float32](nums ...T) T {
	var total T = 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}
