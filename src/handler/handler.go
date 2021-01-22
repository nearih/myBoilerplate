package handler

import (
	"boilerplate/src/handler/example"

	"github.com/labstack/echo"
)

func RegisterHandler(e *echo.Echo) {
	exampleH := example.NewExampleHandler(e)
	exampleH.RegisterExampleHandler()
}
