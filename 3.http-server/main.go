package main

import "fmt"

func sum(nums []int, ch chan int) {
	total := 0

	for _, n := range nums {
		total += n
	}

	ch <- total
}

func main() {
	ch := make(chan int)

	nums := []int{1, 2, 3, 4, 5, 6}

	a := nums[:3]
	b := nums[3:]

	go sum(a, ch)
	go sum(b, ch)

	resA := <-ch
	resB := <-ch

	result := resA + resB

	fmt.Println(result)
}
