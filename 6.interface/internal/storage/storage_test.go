package storage

import (
	"errors"
	"testing"

	"interface/internal/todo"
)

func TestMemoryStorageSaveAndFind(t *testing.T) {
	s := NewMemoryStorage()

	err := s.Save(todo.Todo{
		ID:    1,
		Title: "hello",
		Done:  false,
	})
	if err != nil {
		t.Fatalf("Save() failed, %v", err)
	}

	read, err := s.Find(1)
	if err != nil {
		t.Fatalf("Find() failed, %v", err)
	}

	if read.ID != 1 {
		t.Fatalf("Find() result id is incorrect %d", read.ID)
	}
}

func TestMemoryStorageFindNotFound(t *testing.T) {
	s := NewMemoryStorage()

	_, err := s.Find(1)

	if !errors.Is(err, todo.ErrNotFound{ID: 1}) {
		t.Fatalf("Not expected error, %v", err)
	}
}

func TestMemoryStorageDelete(t *testing.T) {
	s := NewMemoryStorage()
	s.Save(todo.Todo{
		ID:    1,
		Title: "todo",
		Done:  false,
	})
	err := s.Delete(1)
	if err != nil {
		t.Fatalf("Delete() error, %v", err)
	}

	_, err = s.Find(1)
	if !errors.Is(err, todo.ErrNotFound{ID: 1}) {
		t.Fatalf("%d todo must be not found", 1)
	}
}
