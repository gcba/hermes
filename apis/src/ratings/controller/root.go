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

	optionsHeaders := responses.Header{}

	postHeaders := responses.Header{
		ContentType: "application/json; charset=utf-8"
	}

	optionsRatings := responses.Endpoint{
		Method: "OPTIONS",
		Path: "/ratings",
		Headers: optionsHeaders
	}

	postRatings := responses.Endpoint{
		Method: "POST",
		Path: "/ratings",
		Headers: postHeaders
	}

	response := responses.Options{
		Meta:  meta
		Endpoints: []Endpoint{optionsRatings, postRatings}
	}

  	return context.JSON(http.StatusOK, &response)
}