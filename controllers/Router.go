package controllers

import (
	"github.com/kataras/iris"
	TodoController "github.com/paduvi/BasicIrisExample/controllers/todo"
	IndexController "github.com/paduvi/BasicIrisExample/controllers/index"
	MessageController "github.com/paduvi/BasicIrisExample/controllers/message"
	HistoryController "github.com/paduvi/BasicIrisExample/controllers/history"
)

func WithRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	TodoController.EquipRouter(mainRouter)
	IndexController.EquipRouter(mainRouter)
	MessageController.EquipRouter(mainRouter)
	HistoryController.EquipRouter(mainRouter)
}
