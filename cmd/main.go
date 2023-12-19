package main

import (
	"go-htmx-playground/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	helloHandler := handler.HelloHandler{}
	app.GET("/hello", helloHandler.HandleHelloShow)

	app.Start(":8080")
}
