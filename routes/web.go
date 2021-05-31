package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ofarukbicer/echo-boilerplate/handlers/web"
)

func Web(root *echo.Echo) {
	root.GET("/", web.GetIndex).Name = "webGetIndex"
	root.POST("/", web.AddData).Name = "webAddData"
}
