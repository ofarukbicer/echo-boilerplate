package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ofarukbicer/echo-boilerplate/handlers/api"
)

func Api(root *echo.Group) {
	root.GET("/", api.GetIndex).Name = "apiGetIndex"
	root.POST("/", api.AddData).Name = "apiAddData"
}
