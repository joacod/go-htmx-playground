package main

import (
	"go-htmx-playground/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	// Static files
	app.File("/output.css", "css/output.css")
	app.File("/favicon.png", "images/favicon.png")

	// Pages routes
	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandleHomePage)

	helloHandler := handler.HelloHandler{}
	app.GET("/hello", helloHandler.HandleHelloPage)

	// Start server
	app.Start(":8080")
}
