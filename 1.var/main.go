package main

import "fmt"

func main() {
	var f float32 = 42.33
	var i int = int(f)
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
