package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"web/routes"
)

// Template Struct
type Template struct {
	templates *template.Template
}

// Render Function
func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Template Render (html) assets/templates/index.html
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("./assets/templates/*"))}

	// Static Files css,img,js ..
	e.Static("/static", "assets/static")

	// Basic Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Routing Main
	routes.Web(e)

	// Routing Api
	api := e.Group("/api")
	{routes.Api(api)}

	// Debug
	e.Debug = true

	// Run server 8080 port
	e.Start(":8080")
}