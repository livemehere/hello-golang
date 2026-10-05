package storage

import todo "interface/internal/todo"

type Storage interface {
	Save(todo todo.Todo) error
	Find(id int) (todo.Todo, error)
	Delete(id int) error
}
