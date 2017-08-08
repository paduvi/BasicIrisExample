package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/handlers/index"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipIndexRouter(app router.Party) {
	party := app.Party("/")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"SubIndex",
		"GET",
		"/{name}",
		SubIndex,
	},
}
