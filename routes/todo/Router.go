package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/handlers/todo"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipTodoRouter(app router.Party) {
	party := app.Party("/todos")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/{todoId:int min(1)}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/",
		TodoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/{todoId:int min(1)}",
		TodoDelete,
	},
}
