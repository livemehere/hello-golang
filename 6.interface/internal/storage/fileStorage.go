package storage

import (
	"encoding/json"
	"fmt"
	"os"

	todo "interface/internal/todo"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}

func (s *FileStorage) Find(id int) (todo.Todo, error) {
	todos, err := s.load()
	if err != nil {
		return todo.Todo{}, err
	}

	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return todo.Todo{}, &todo.ErrNotFound{
		ID: id,
	}
}

func (s *FileStorage) Save(todo todo.Todo) error {
	todos, err := s.load()
	if err != nil {
		return err
	}

	found := false

	for i := range todos {
		if todos[i].ID == todo.ID {
			todos[i] = todo
			found = true
			break
		}
	}

	if !found {
		todos = append(todos, todo)
	}
	return s.saveAll(todos)
}

func (s *FileStorage) Delete(id int) error {
	todos, err := s.load()
	if err != nil {
		return err
	}

	index := -1

	for i, todo := range todos {
		if todo.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return &todo.ErrNotFound{
			ID: id,
		}
	}

	todos = append(
		todos[:index],
		todos[index+1:]...,
	)

	return s.saveAll(todos)
}

func (s *FileStorage) load() ([]todo.Todo, error) {
	file, err := os.Open(s.path)
	if err != nil {
		// if os.IsNotExist(err) {
		// 	if err := os.WriteFile(s.path, []byte("[]"), 0o644); err != nil {
		// 		return nil, err
		// 	}
		// 	return []todo.Todo{}, nil
		// }

		return nil, fmt.Errorf("file not found %s : %w", s.path, err)
	}

	defer file.Close()

	var todos []todo.Todo

	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *FileStorage) saveAll(todos []todo.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0o644)
}
