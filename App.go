package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	"github.com/paduvi/BasicIrisExample/controllers"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	//app.Use(middlewares.Logger) // uncomment to see log

	controllers.WithRouter(app)

	app.Run(iris.Addr(os.Getenv("Address")), iris.WithCharset("UTF-8"))
}
