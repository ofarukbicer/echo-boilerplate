package web

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"github.com/ofarukbicer/echo-boilerplate/database"
)

func GetIndex(c echo.Context) error {
	db := database.Client("deneme","deneme")
	data := db.Find(bson.M{})
	return c.Render(
		http.StatusOK,
		"index.html",
		echo.Map{
			"data": data,
		},
	)
}

func AddData(c echo.Context) error {
	db := database.Client("deneme","deneme")
	db.InsertOne(bson.D{
		{ Key: "name", Value: "Normal" },
	})
	return c.Redirect(http.StatusSeeOther, "/")
}
