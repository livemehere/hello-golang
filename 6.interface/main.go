package main

import (
	"errors"
	"fmt"

	"interface/internal/storage"
	"interface/internal/todo"
)

func main() {
	var s storage.Storage

	s = storage.NewFileStorage("non.josn")
	s.Save(todo.Todo{
		ID:    1,
		Title: "hello world",
	})

	{
		fmt.Println("FINDING...")
		todo, err := s.Find(1)
		fmt.Println(todo, err)
	}

	{
		fmt.Println("DELETING...")
		s.Delete(1)
	}

	var notfoundErr todo.ErrNotFound
	_, err := s.Find(1)
	if errors.As(err, &notfoundErr) {
		fmt.Printf("NOT FOUND!!! ID : %d\n", notfoundErr.ID)
	}
}
