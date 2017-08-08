package routes

import (
	"github.com/kataras/iris"
	. "github.com/paduvi/BasicIrisExample/routes/todo"
	. "github.com/paduvi/BasicIrisExample/routes/index"
	. "github.com/paduvi/BasicIrisExample/routes/message"
)

func EquipRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	EquipIndexRouter(mainRouter)
	EquipTodoRouter(mainRouter)
	EquipMessageRouter(mainRouter)
}
