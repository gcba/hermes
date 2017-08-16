package controller

import (
	"hermes/ratings/responses"

	"github.com/labstack/echo"
)

func OptionsRoot(context echo.Context) error {
	endpoints := []responses.Endpoint{
		responses.Endpoints["OptionsRatings"],
		responses.Endpoints["PostRatings"]}

	context.Response().Header().Set(echo.HeaderAllow, "OPTIONS")

	return responses.OptionsResponse(endpoints, context)
}
