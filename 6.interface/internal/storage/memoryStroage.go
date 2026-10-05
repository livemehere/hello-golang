package storage

import (
	"fmt"

	"interface/internal/todo"
)

type MemoryStorage struct {
	todos map[int]todo.Todo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[int]todo.Todo),
	}
}

func (s *MemoryStorage) Save(todo todo.Todo) error {
	s.todos[todo.ID] = todo

	return nil
}

func (s *MemoryStorage) Find(id int) (todo.Todo, error) {
	found, ok := s.todos[id]

	if !ok {
		return todo.Todo{}, fmt.Errorf("failed to find todo %d, %w", id, todo.ErrNotFound{
			ID: id,
		})
	}

	return found, nil
}

func (s *MemoryStorage) Delete(id int) error {
	_, ok := s.todos[id]
	if !ok {
		return todo.ErrNotFound{ID: id}
	}

	delete(s.todos, id)

	return nil
}
