package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/models"
	. "github.com/paduvi/BasicIrisExample/handlers/message"
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
		TodoIndex,
	},
	Route{
		"MessagePing",
		"GET",
		"/ping",
		MessagePing,
	},
	Route{
		"MessageShow",
		"GET",
		"/{messageId:int min(1)}",
		TodoShow,
	},
	Route{
		"MessageCreate",
		"POST",
		"/",
		TodoCreate,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/{messageId:int min(1)}",
		TodoDelete,
	},
}
