package main

import (
	"./usuario"

	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {

	e.POST("/api/v1/users", usuario.Create)
	e.GET("/api/v1/users", usuario.GetAll)
	e.GET("/api/v1/users/:email", usuario.GetByEmail)
	e.GET("api/v1/users-paginate", usuario.GetAllPaginate)

}
