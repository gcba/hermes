package controller

import (
	"net/http"

	"ratings/parser"

	"github.com/labstack/echo"
)

func CreateRating(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, request)
}
