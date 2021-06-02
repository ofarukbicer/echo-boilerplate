package api

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"web/database"
)

// GetIndex GET - /api - Handler
func GetIndex(c echo.Context) error {
	db := database.Client("deneme", "deneme")
	data := db.Find(bson.M{})
	return c.JSON(
		http.StatusOK,
		data,
	)
}

// AddData POST - /api - Handler
func AddData(c echo.Context) error {
	db := database.Client("deneme","deneme")
	db.InsertOne(bson.D{
		{ Key: "name", Value: "Api" },
	})
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"text": "success",
			"code": 200,
		},
	)
}