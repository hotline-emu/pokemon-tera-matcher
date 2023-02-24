package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommonResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

func BadRequest(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		CommonResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(
		http.StatusInternalServerError,
		CommonResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}
