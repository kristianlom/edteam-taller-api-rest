package main

import (
	"./zapato"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create)
}
