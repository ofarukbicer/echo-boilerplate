package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"web/routes"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("./assets/templates/*"))}
	e.Static("/static", "assets/static")
	routes.Web(e)
	api := e.Group("/api")
	{routes.Api(api)}
	e.Debug = true
	e.Start(":8080")
}