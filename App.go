package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	. "github.com/paduvi/BasicIrisExample/routes"
	. "github.com/paduvi/BasicIrisExample/httputils"
	"github.com/paduvi/BasicIrisExample/config"
)

var dispatcher *Dispatcher

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	//app.Use(middlewares.Logger) // uncomment to see log

	EquipRouter(app)

	dispatcher = NewDispatcher(config.MaxWorker)
	dispatcher.Run()

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
