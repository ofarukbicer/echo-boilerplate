package routes

import (
	"github.com/labstack/echo/v4"
	"web/handlers/api"
)

func Api(root *echo.Group) {
	// GET - /api
	root.GET("", api.GetIndex).Name = "apiGetIndex"
	// POST - /api
	root.POST("", api.AddData).Name = "apiAddData"
}