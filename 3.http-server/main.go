package main

import (
	"fmt"
)

type UserID int

type Person interface {
	SetName(name string)
}

type User struct {
	Id   UserID
	Name string
	age  int
}

func (u User) hello() {
	fmt.Println(u.Id, u.Name, u.age)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func DoSomthing(p Person) {
	p.SetName("DODODO")
}

func main() {
	me := User{
		Id:   10,
		Name: "Unknown",
		age:  99,
	}

	var n UserID = 88
	me.Id = n

	//me.SetName("Kong")
	DoSomthing(&me)
	me.hello()
}
