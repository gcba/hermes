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

	optionsRatings := responses.Endpoint{
		Method: "OPTIONS",
		Path: "/ratings"
	}

	postRatings := responses.Endpoint{
		Method: "POST",
		Path: "/ratings",
		Headers: &responses.Header{
			ContentType: "application/json; charset=utf-8"
		}
	}

	response := responses.Options{
		Meta:  meta
		Endpoints: []responses.Endpoint{optionsRatings, postRatings}
	}

  	return context.JSON(http.StatusOK, &response)
}