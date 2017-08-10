package todo

import (
	"github.com/kataras/iris/core/router"
	. "github.com/paduvi/BasicIrisExample/models"
)

func EquipRouter(app router.Party) {
	party := app.Party("/histories")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"GET",
		"/",
		ListViewer,
	},
	Route{
		"GET",
		"/{itemId:int}",
		ListViewerByItemId,
	},
	Route{
		"GET",
		"/{itemId:int}/{userId:int}",
		ViewItemByUserId,
	},
	Route{
		"GET",
		"/user/{userId:int}",
		ShowUserHistory,
	},
}
