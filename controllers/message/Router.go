package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipRouter(app router.Party) {
	party := app.Party("/messages")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"GET",
		"/",
		MessageIndex,
	},
	Route{
		"GET",
		"/{messageId:int min(1)}",
		MessageShow,
	},
	Route{
		"POST",
		"/",
		MessageCreate,
	},
	Route{
		"DELETE",
		"/{messageId:int min(1)}",
		MessageDelete,
	},
}
