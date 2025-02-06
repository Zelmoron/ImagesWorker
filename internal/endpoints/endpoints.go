package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
	}
	Endpoints struct {
		services Services
	}
)

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}

func (e *Endpoints) Render(c echo.Context) error {

	return c.Render(http.StatusOK, "hello", nil)

}
