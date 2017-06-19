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
	meta := responses.Meta{
		Code: 200,
		Message: "Request completed successfully"
	}

	response := responses.Options{
		Meta:  meta
		Endpoints: []responses.Endpoint{
			responses.OptionsRatings,
			responses.PostRatings
		}
	}

  	return context.JSON(http.StatusOK, &response)
}