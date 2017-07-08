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

	if !context.Response().Committed {
		return ErrorsResponse(status, []string{errorMessage}, context)
	}

	return nil
}

func ErrorsResponse(status int, errorList []string, context echo.Context) error {
	if len(errorList) == 0 {
		response := BasicError{Meta: metas[status]}

		return context.JSON(status, &response)
	}

	errorMessages := ""

	for _, message := range errorList {
		errorMessages += message + "\n"
	}

	context.Logger().Error(errorMessages)

	if !context.Response().Committed {
		response := Error{
			Meta:   metas[status],
			Errors: errorList}

		return context.JSON(status, &response)
	}

	return nil
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
