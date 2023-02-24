package main

import (
	"net/http"
	"pokemon-tera-matcher/configs"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MONGO_URI = "mongodb://localhost:27017"

func main() {
	e := echo.New()
	configs.ConnectDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type PokemonType struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"text"`
	Stengths   []string           `bson:"array"`
	Weaknesses []string           `bson:"array"`
	ImmuneTo   []string           `bson:"array"`
	CannotHit  []string           `bson:"array"`
}
