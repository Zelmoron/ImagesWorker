package app

import (
	"ImageWorkr/internal/endpoints"
	"ImageWorkr/internal/services"

	"github.com/labstack/echo/v4"
)

type App struct {
	app       *echo.Echo
	services  *services.Services
	endpoints *endpoints.Endpoints
}

func New() *App {
	a := &App{}

	a.app = echo.New()
	a.services = services.New()
	a.endpoints = endpoints.New(a.services)

	return a
}

func (a *App) routers() {

}

func (a *App) Run() {
	a.app.Start(":8080")
}
