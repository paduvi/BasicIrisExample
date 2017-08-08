package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipMessageRouter(app router.Party) {
	party := app.Party("/messages")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"MessageIndex",
		"GET",
		"/",
		MessageIndex,
	},
	Route{
		"MessageShow",
		"GET",
		"/{messageId:int min(1)}",
		MessageShow,
	},
	Route{
		"MessageCreate",
		"POST",
		"/",
		MessageCreate,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/{messageId:int min(1)}",
		MessageDelete,
	},
}
