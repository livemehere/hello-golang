package main

import (
	"fmt"
)

type User struct {
	Name string
	age  int
}

func (u User) hello() {
	fmt.Println(u.Name, u.age)
}

func main() {
	me := User{
		Name: "안녕",
		age:  99,
	}

	me.hello()
}
