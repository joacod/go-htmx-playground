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
	app.File("/htmx.min.js", "js/htmx.min.js")

	// Pages routes
	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandleHomePage)

	// API routes
	helloHandler := handler.HelloHandler{}
	app.GET("/api/hello", helloHandler.HandleHelloComponent)

	// Start server
	app.Start(":8080")
}
