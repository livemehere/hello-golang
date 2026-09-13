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
}
