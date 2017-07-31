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

func ErrorResponse(status int, errorList []string, context echo.Context) error {
	if !context.Response().Committed {
		if len(errorList) == 0 {
			response := BasicError{Meta: metas[status]}

			return context.JSON(status, &response)
		}

		response := Error{
			Meta:   metas[status],
			Errors: errorList}

		return context.JSON(status, &response)
	}

	return nil
}

func ErrorHandler(err error, context echo.Context) {
	status := http.StatusInternalServerError
	messages := []string{err.Error()}

	if echoHTTPError, ok := err.(*echo.HTTPError); ok {
		status = echoHTTPError.Code
		messages = echoHTTPError.Message.([]string)
	}

	if !context.Response().Committed {
		ErrorResponse(status, messages, context)
	}
}
