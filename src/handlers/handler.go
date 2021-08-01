package handlers

import (
	"boilerplate/src/handlers/example"

	"github.com/labstack/echo"
)

func RegisterHandler(e *echo.Echo) {
	exampleH := example.NewExampleHandler(e)
	exampleH.RegisterExampleHandler()
}
