package usuario

import (
	"../respuesta"
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context) error {
	u := &Model{}
	err := c.Bind(u)
	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "E005",
				Contenido: "Mal objeto",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}
	d := storage.Login(u.Email, u.Password)
	d.Password = ""
	if d == nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "L001",
				Contenido: "Usuario o password incorrectos",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}
	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "O001",
			Contenido: "Logueado ok",
		},
		Data: d,
	}
	return c.JSON(http.StatusOK, r)
}
