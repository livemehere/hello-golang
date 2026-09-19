package main

import "fmt"

func main() {
	composed := "한"
	decomposed := "\u1112\u1161\u11AB"

	fmt.Println(composed)
	fmt.Println(decomposed)

	fmt.Println(composed == decomposed) // false
}
