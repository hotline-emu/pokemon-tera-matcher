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
	"go.mongodb.org/mongo-driver/bson"
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
		return responses.BadRequest(c, err)
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&pokemonType); validationErr != nil {
		return responses.BadRequest(c, validationErr)
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
		return responses.InternalServerError(c, err)
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

func GetAllPokemonTypes(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var PokemonTypes []models.PokemonType
	defer cancel()

	results, err := pokemonTypeCollection.Find(ctx, bson.M{})

	if err != nil {
		return responses.InternalServerError(c, err)
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePokemonType models.PokemonType
		if err = results.Decode(&singlePokemonType); err != nil {
			return responses.InternalServerError(c, err)
		}

		PokemonTypes = append(PokemonTypes, singlePokemonType)
	}

	return c.JSON(
		http.StatusOK,
		responses.CommonResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &echo.Map{"data": PokemonTypes},
		},
	)
}

func GetAPokemonType(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	pokemonTypeId := c.Param("id")
	var PokemonType models.PokemonType
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(pokemonTypeId)

	err := pokemonTypeCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&PokemonType)
	if err != nil {
		return responses.InternalServerError(c, err)
	}

	return c.JSON(
		http.StatusOK,
		responses.CommonResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &echo.Map{"data": PokemonType},
		},
	)
}
