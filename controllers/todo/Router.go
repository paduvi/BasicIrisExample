package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipRouter(app router.Party) {
	party := app.Party("/todos")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"GET",
		"/",
		TodoIndex,
	},
	Route{
		"GET",
		"/{todoId:int min(1)}",
		TodoShow,
	},
	Route{
		"POST",
		"/",
		TodoCreate,
	},
	Route{
		"DELETE",
		"/{todoId:int min(1)}",
		TodoDelete,
	},
}
