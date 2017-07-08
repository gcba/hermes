package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	Error struct {
		Meta   meta     `json:"meta"`
		Errors []string `json:"errors"`
	}

	BasicError struct {
		Meta meta `json:"meta"`
	}
)

func ErrorResponse(status int, errorMessage string, context echo.Context) error {
	if errorMessage == "" {
		return ErrorsResponse(status, []string{}, context)
	}

	context.Logger().Error(errorMessage)

	return ErrorsResponse(status, []string{errorMessage}, context)
}

func ErrorsResponse(status int, errors []string, context echo.Context) error {
	if len(errors) == 0 {
		response := BasicError{Meta: metas[status]}

		return context.JSON(status, &response)
	}

	for _, errorMessage := range errors {
		context.Logger().Error(errorMessage)
	}

	response := Error{
		Meta:   metas[status],
		Errors: errors}

	return context.JSON(status, &response)
}

func ErrorHandler(err error, context echo.Context) {
	status := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
	} else if context.Echo().Debug {
		context.Logger().Error(err.Error())
	}

	if !context.Response().Committed {
		ErrorsResponse(status, []string{}, context)
	}
}
