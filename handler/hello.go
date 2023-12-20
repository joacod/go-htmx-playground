package handler

import (
	"go-htmx-playground/view/hello"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandleHelloPage(c echo.Context) error {
	return render(c, hello.Page())
}
