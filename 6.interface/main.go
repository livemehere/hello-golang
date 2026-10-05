package main

import (
	"errors"
	"fmt"

	"interface/internal/storage"
	"interface/internal/todo"
)

type Storage interface {
	Save(todo todo.Todo) error
	Find(id int) (todo.Todo, error)
	Delete(id int) error
}

func main() {
	var s Storage

	s = storage.NewMemoryStorage()
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
