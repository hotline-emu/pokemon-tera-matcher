package routes

import (
	"pokemon-tera-matcher/controllers"

	"github.com/labstack/echo/v4"
)

func PokemonTypeRoute(e *echo.Echo) {
	e.POST("/pokemon-type", controllers.CreatePokemonType)
	e.GET("/pokemon-type", controllers.GetAllPokemonTypes)
	e.GET("/pokemon-type/:id", controllers.GetAPokemonType)
}
