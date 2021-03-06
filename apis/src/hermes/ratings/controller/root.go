package controller

import (
	"hermes/ratings/responses"

	"github.com/labstack/echo"
)

func OptionsRoot(context echo.Context) error {
	endpoints := []responses.Endpoint{
		responses.Endpoints["OptionsRatings"],
		responses.Endpoints["PostRatings"]}

	return responses.OptionsResponse(endpoints, "OPTIONS", context)
}
