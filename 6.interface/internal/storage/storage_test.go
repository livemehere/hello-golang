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
	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name: "delete existing todo",
			id:   1,
		},
		{
			name:    "delete missing todo",
			id:      999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			s := NewMemoryStorage()
			s.Save(todo.Todo{
				ID:    1,
				Title: "sample",
				Done:  false,
			})

			// expect
			err := s.Delete(tt.id)

			if tt.wantErr && err == nil {
				t.Fatalf("expect error")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		})
	}
}
