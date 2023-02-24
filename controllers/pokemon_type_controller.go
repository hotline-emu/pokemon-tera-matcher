package controllers

import (
	"context"
	"net/http"
	"pokemon-tera-matcher/configs"
	"pokemon-tera-matcher/models"
	"pokemon-tera-matcher/responses"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var pokemonTypeCollection *mongo.Collection = configs.GetCollection(configs.DB, "pokemonTypes")
var validate = validator.New()

func CreatePokemonType(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pokemonType models.PokemonType
	defer cancel()

	//validate the request body
	if err := c.Bind(&pokemonType); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			responses.CommonResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			},
		)
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&pokemonType); validationErr != nil {
		return c.JSON(
			http.StatusBadRequest,
			responses.CommonResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": validationErr.Error()},
			},
		)
	}

	newPokemonType := models.PokemonType{
		Id:         primitive.NewObjectID(),
		Name:       pokemonType.Name,
		Stengths:   pokemonType.Stengths,
		Weaknesses: pokemonType.Weaknesses,
		ImmuneTo:   pokemonType.ImmuneTo,
		CannotHit:  pokemonType.CannotHit,
	}

	result, err := pokemonTypeCollection.InsertOne(ctx, newPokemonType)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			responses.CommonResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		responses.CommonResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &echo.Map{"data": result},
		},
	)
}
