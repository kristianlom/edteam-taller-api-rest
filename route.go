package main

import (
	"./usuario"
	"./zapato"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create)
	e.POST("/usuarios", usuario.Login)
}
