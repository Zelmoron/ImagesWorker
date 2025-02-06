package app

import (
	"ImageWorkr/internal/endpoints"
	"ImageWorkr/internal/services"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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

	a.routers()

	return a
}

func (a *App) routers() {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	a.app.Renderer = t
	a.app.Debug = true
	a.app.Static("/static", "static")
	a.app.Use(middleware.Logger())
	a.app.GET("/", a.endpoints.Render)
	a.app.POST("/register/image", a.endpoints.ImageWork)
}

func (a *App) Run() {
	a.app.Start(":8080")
}
