package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	"github.com/paduvi/BasicIrisExample/controllers"
	"os"
	"github.com/paduvi/BasicIrisExample/httputils"
	"strconv"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	//app.Use(middlewares.Logger) // uncomment to see log

	controllers.WithRouter(app)

	MaxWorker, _ := strconv.Atoi(os.Getenv("MaxWorker"))
	httputils.NewDispatcher(MaxWorker).Run()

	app.Run(iris.Addr(os.Getenv("Address")), iris.WithCharset("UTF-8"))
}
