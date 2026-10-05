package todo

import "fmt"

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type ErrNotFound struct {
	ID int
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("todo %d not found", e.ID)
}
