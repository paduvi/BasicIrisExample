package actions

import "github.com/paduvi/BasicIrisExample/mockrepo"
import . "github.com/paduvi/BasicIrisExample/models"

func FindTodo(id int, done chan Todo) {
	// return empty Todo if not found
	done <- mockrepo.FindTodo(id)
}

func ListTodo(done chan Todos) {
	done <- mockrepo.ListTodo()
}

func CreateTodo(t Todo, done chan Todo) {
	done <- mockrepo.CreateTodo(t)
}

func DestroyTodo(id int, done chan error) {
	done <- mockrepo.DestroyTodo(id)
}
