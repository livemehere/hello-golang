package main

import f "fmt"

func main() {
	name := "한글"
	runes := []rune(name)

	f.Println("len is ", len(runes))

	for i, r := range name {
		f.Printf("%d : %c\n", i, r)
	}
}
