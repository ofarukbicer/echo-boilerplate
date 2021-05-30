package routes

import (
	"github.com/labstack/echo/v4"
	"web/handlers/api"
)

func Api(root *echo.Group) {
	root.GET("/", api.GetIndex).Name = "apiGetIndex"
	root.POST("/", api.AddData).Name = "apiAddData"
}