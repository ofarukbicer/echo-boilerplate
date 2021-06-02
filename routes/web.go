package routes

import (
	"github.com/labstack/echo/v4"
	"web/handlers/web"
)

func Web(root *echo.Echo) {
	// GET - /
	root.GET("/", web.GetIndex).Name = "webGetIndex"
	// POST - /
	root.POST("/", web.AddData).Name = "webAddData"
}