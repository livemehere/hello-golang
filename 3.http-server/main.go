package main

import "fmt"

func main() {
	fmt.Println("hello world")

	n := 10
	var p *int

	p = &n
	*p = 99

	fmt.Println(*p, n)
}
