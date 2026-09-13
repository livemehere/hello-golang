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

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	me := User{
		Name: "Unknown",
		age:  99,
	}

	me.SetName("Kong")
	me.hello()
}
