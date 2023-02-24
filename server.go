package main

import (
	"pokemon-tera-matcher/configs"
	"pokemon-tera-matcher/routes"

	"github.com/labstack/echo/v4"
)

const MONGO_URI = "mongodb://localhost:27017"

func main() {
	e := echo.New()
	configs.ConnectDB()

	routes.PokemonTypeRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
