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

	switch v := x.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	}
}
