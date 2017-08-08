package controllers

import (
	"github.com/kataras/iris"
	. "github.com/paduvi/BasicIrisExample/controllers/todo"
	. "github.com/paduvi/BasicIrisExample/controllers/index"
	. "github.com/paduvi/BasicIrisExample/controllers/message"
)

func WithRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	EquipIndexRouter(mainRouter)
	EquipTodoRouter(mainRouter)
	EquipMessageRouter(mainRouter)
}
