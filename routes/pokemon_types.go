package routes

import (
	"pokemon-tera-matcher/controllers"

	"github.com/labstack/echo/v4"
)

func PokemonTypeRoute(e *echo.Echo) {
	e.POST("/pokemon-type", controllers.CreatePokemonType)
}
