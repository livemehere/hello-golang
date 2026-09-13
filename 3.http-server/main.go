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

	result, err := math.Divide(10, 2)

	if err != nil && err.Error() == "division by zero" {
		fmt.Println("Error!!!!", err)
	} else {
		fmt.Println(result, err)
	}

	// switch

	a := 22

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("nothing")
	}
}
