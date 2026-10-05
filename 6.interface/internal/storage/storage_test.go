package storage

import (
	"testing"

	todo "interface/internal/todo"
)

func TestMemoryStorage(t *testing.T) {
	s := NewMemoryStorage()

	err := s.Save(todo.Todo{
		ID:    1,
		Title: "hello",
		Done:  false,
	})
	if err != nil {
		t.Fatalf("Save() err = %v", err)
	}

	_, err = s.Find(1)
	if err != nil {
		t.Fatalf("Find() err = %v", err)
	}
}
