package handler

import (
	"go-htmx-playground/api/hello"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandleHelloComponent(c echo.Context) error {
	return render(c, hello.Component())
}
