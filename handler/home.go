package handler

import (
	"go-htmx-playground/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomePage(c echo.Context) error {
	return render(c, home.Page())
}
