package main

import "fmt"

func main() {
	const f float32 = 42.33
	var i int = 1
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
