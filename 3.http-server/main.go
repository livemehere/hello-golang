package main

import (
	"fmt"
)

func main() {
	var x any = 19

	s, ok := x.(string)

	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("failed!", ok)
	}

	defer fmt.Println("string:")
	defer fmt.Println("int:")

	f := func(x int) int {
		return x + 10
	}

	fmt.Println(f(1))
}
