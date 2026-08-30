package main

import "fmt"

var a, b, c bool = true, true, true

// k := 190 (disabled outside of func)

func main() {
	var i int = 99
	k := 190 // enable inside of func
	fmt.Println(a, b, c, i, k)
}
