package controller

import (
	"ratings/responses"

	"github.com/labstack/echo"
)

func OptionsRoot(context echo.Context) error {
	endpoints := []responses.Endpoint{
		responses.Endpoints["OptionsRatings"],
		responses.Endpoints["PostRatings"]}

	return responses.OptionsResponse(endpoints, context)
}
