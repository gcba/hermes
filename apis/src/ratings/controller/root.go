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

	optionsRatings := responses.Method{
		Verb: "OPTIONS",
		Endpoint: "/ratings",
		Headers: optionsHeaders
	}

	postRatings := responses.Method{
		Verb: "POST",
		Endpoint: "/ratings",
		Headers: postHeaders
	}

	response := responses.Options{
		Meta:  meta
		Methods: []Method{optionsRatings, postRatings}
	}

  	return c.JSON(http.StatusOK, &response)
}