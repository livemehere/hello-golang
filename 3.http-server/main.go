package main

import f "fmt"

func sum[T int | float32](nums ...T) T {
	var total T = 0
	for _, n := range nums {
		total += n
	}
	return total
}

func init() {
	f.Println("init")
}

func main() {
	f.Println("main")
	nums := []int{1, 2, 3, 4}
	f.Println(sum(nums...))
}
