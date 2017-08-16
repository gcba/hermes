package responses

import (
	"net/http"
	"strings"

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

func ErrorResponse(status int, singleError string, context echo.Context) error {
	if !context.Response().Committed {
		if len(strings.Trim(singleError, " ")) == 0 {
			return ErrorsResponse(status, []string{}, context)
		}

		return ErrorsResponse(status, []string{singleError}, context)
	}

	return nil
}

func ErrorsResponse(status int, errorList []string, context echo.Context) error {
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
	var messages []string

	isString := false
	isArray := false
	status := http.StatusInternalServerError

	if echoHTTPError, ok := err.(*echo.HTTPError); ok {
		status = echoHTTPError.Code
		messages, isArray = echoHTTPError.Message.([]string)
		_, isString = echoHTTPError.Message.(string)
	}

	if !context.Response().Committed {
		if isArray {
			ErrorsResponse(status, messages, context)
		} else if isString {
			ErrorResponse(status, "", context)
		} else {
			ErrorResponse(status, err.Error(), context)
		}
	}
}
