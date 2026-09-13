package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	for i, v := range nums {
		fmt.Println(i, v)
		nums[i] = v + 10
	}

	fmt.Println("----------------------")

	for i, v := range nums2 {
		fmt.Println(i, v)
	}

	fmt.Println("len is", len(nums))
}
