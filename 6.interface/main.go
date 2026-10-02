package main

import (
	"errors"
	"fmt"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Storage interface {
	Save(todo Todo) error
	Find(id int) (Todo, error)
	Delete(id int) error
}

type MemoryStorage struct {
	todos map[int]Todo
}

var ErrNotFound = errors.New("todo not found")

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[int]Todo),
	}
}

func (s *MemoryStorage) Save(todo Todo) error {
	s.todos[todo.ID] = todo

	return nil
}

func (s *MemoryStorage) Find(id int) (Todo, error) {
	found, ok := s.todos[id]

	if !ok {
		return Todo{}, fmt.Errorf("failed to find todo %d, %w", id, ErrNotFound)
	}

	return found, nil
}

func (s *MemoryStorage) Delete(id int) error {
	_, ok := s.todos[id]
	if !ok {
		return ErrNotFound
	}

	delete(s.todos, id)

	return nil
}

func main() {
	storage := NewMemoryStorage()
	storage.Save(Todo{
		ID:    1,
		Title: "hello world",
	})

	{
		fmt.Println("FINDING...")
		todo, err := storage.Find(1)
		fmt.Println(todo, err)
	}

	{
		fmt.Println("DELETING...")
		storage.Delete(1)
	}

	_, err := storage.Find(1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("NOT FOUND!!!", err)
	}
}
