package todo

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	. "github.com/paduvi/BasicIrisExample/models"
	TodoAction "github.com/paduvi/BasicIrisExample/actions/todo"

	"io/ioutil"
	"io"
	"github.com/kataras/iris"
)

func TodoIndex(ctx context.Context) {
	ctx.ContentType("application/json")
	ctx.StatusCode(iris.StatusOK)
	done := make(chan Todos)
	go TodoAction.ListTodo(done)

	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func TodoShow(ctx context.Context) {
	todoId, err := ctx.Params().GetInt("todoId")
	if err != nil {
		ctx.Writef("error while trying to parse userid parameter," +
			"this will never happen if :int is being used because if it's not integer it will fire Not Found automatically.")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan Todo)
	go TodoAction.FindTodo(todoId, done)
	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func TodoCreate(ctx context.Context) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(ctx.Request().Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := ctx.Request().Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		ctx.ContentType("application/json")
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		if _, err := ctx.JSON(err); err != nil {
			panic(err)
		}
	}

	done := make(chan Todo)
	go TodoAction.CreateTodo(todo, done)
	ctx.ContentType("application/json")
	ctx.StatusCode(iris.StatusCreated)
	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func TodoDelete(ctx context.Context) {
	todoId, err := ctx.Params().GetInt("todoId")
	if err != nil {
		ctx.Writef("error while trying to parse userid parameter," +
			"this will never happen if :int is being used because if it's not integer it will fire Not Found automatically.")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan error)
	go TodoAction.DestroyTodo(todoId, done)
	if err := <-done; err != nil {
		panic(err)
	}
	ctx.Text("Destroy Todo Successfully.")
}
