package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	"github.com/paduvi/BasicIrisExample/controllers"
	"os"
	"github.com/paduvi/BasicIrisExample/httputils"
	"strconv"
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/paduvi/BasicIrisExample/redisutils"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	if os.Getenv("LogLevel") == "DEBUG" {
		app.Use(middlewares.Logger)
	}

	controllers.WithRouter(app)

	MaxWorker, _ := strconv.Atoi(os.Getenv("MaxWorker"))
	httputils.NewDispatcher(MaxWorker).Run()
	redisutils.NewDispatcher(MaxWorker).Run()

	app.Run(iris.Addr(os.Getenv("Address")), iris.WithCharset("UTF-8"))
}
