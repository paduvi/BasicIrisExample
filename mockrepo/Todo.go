package mockrepo

import (
	"fmt"
	. "github.com/paduvi/BasicIrisExample/models"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	go CreateTodo(Todo{Name: "Write presentation"})
	go CreateTodo(Todo{Name: "Host meetup"})
}

func FindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func ListTodo() Todos {
	return todos
}

func CreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func DestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
