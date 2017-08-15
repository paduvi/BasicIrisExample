package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	"github.com/paduvi/BasicIrisExample/controllers"
	"os"
	_ "github.com/jpfuentes2/go-env/autoload"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	if os.Getenv("LogLevel") == "DEBUG" {
		app.Use(middlewares.Logger)
	}

	controllers.WithRouter(app)

	app.Run(iris.Addr(os.Getenv("Address")), iris.WithCharset("UTF-8"))
}
