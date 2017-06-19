package controller

import (
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func OptionsRoot(context echo.Context) error {
	response := responses.Options{
		Meta:  responses.Meta200,
		Endpoints: []responses.Endpoint{
			responses.OptionsRatings,
			responses.PostRatings
		}
	}

  	return context.JSON(http.StatusOK, &response)
}