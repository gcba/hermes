package parser

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(factory func() interface{}, escape func(interface{}) error, context echo.Context) (interface{}, error) {
	request := factory()

	conform.Strings(request)

	if err := escape(request); err != nil {
		return request, err
	}

	if err := bind(request, context); err != nil {
		return request, err
	}

	if err := validate(request, context); err != nil {
		return request, err
	}

	return request, nil
}

func bind(request interface{}, context echo.Context) error {
	if err := context.Bind(request); err != nil {
		errorDescription := err.Error()
		errorMessage := fmt.Sprintf("Error parsing request: %s", errorDescription)
		errorCode := http.StatusBadRequest

		if httpError, ok := err.(*echo.HTTPError); ok {
			if value, isString := httpError.Message.(string); isString {
				errorMessage = value
				errorCode = httpError.Code
			}
		}

		context.Logger().Error("Error binding request: ", errorDescription)

		return echo.NewHTTPError(errorCode, []string{errorMessage})
	}

	return nil
}

func validate(request interface{}, context echo.Context) error {
	if errs := context.Validate(request); errs != nil {
		var errorList []string
		var errorMessage = "Error validating request: "

		if _, ok := errs.(*validator.InvalidValidationError); ok {
			context.Logger().Error(errorMessage, errs.Error())

			return echo.NewHTTPError(http.StatusUnprocessableEntity, []string{errs.Error()})
		}

		for _, err := range errs.(validator.ValidationErrors) {
			errorDescription := err.(error).Error()
			errorList = append(errorList, errorDescription)

			context.Logger().Error(errorMessage, errorDescription)
		}

		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorList)
	}

	return nil
}
