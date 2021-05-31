package routes

import (
	"github.com/labstack/echo/v4"
	"web/handlers/web"
)

func Web(root *echo.Echo) {
	root.GET("/", web.GetIndex).Name = "webGetIndex"
	root.POST("/", web.AddData).Name = "webAddData"
}
