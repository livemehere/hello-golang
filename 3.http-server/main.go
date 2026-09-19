package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "kong"},
		Breed:  "corgi",
	}

	fmt.Println(d)
}
