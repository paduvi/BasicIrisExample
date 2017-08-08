package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/middlewares"
	"github.com/paduvi/BasicIrisExample/controllers"
	"github.com/joho/godotenv"
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

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
