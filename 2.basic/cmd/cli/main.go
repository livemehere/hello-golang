package main

import (
	"basic/internal/auth"
	"fmt"
)

func main() {
	name := auth.GetAuth()

	fmt.Println("Hello world", name, auth.Sum(1, 2))
}
