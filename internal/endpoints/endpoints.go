package endpoints

import "github.com/labstack/echo/v4"

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

func Render(c *echo.Context) error {
	return nil

}
