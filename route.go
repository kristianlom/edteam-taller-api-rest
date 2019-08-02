package main

import (
	"./usuario"
	"./zapato"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create, usuario.ValidateJWT)
	e.POST("/usuarios", usuario.Login)
}
