package parser

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/microcosm-cc/bluemonday"
)

type (
	// Request holds the mapped fields from the request's JSON body
	Request struct {
		Query         string                 `json:"query" validate:"required,gte=10,lte=5000" conform:"trim"`
		OperationName string                 `json:"operationName" validate:"required,gte=3,lte=15,alphanum,excludesall= " conform:"trim,lower"`
		Variables     map[string]interface{} `json:"variables" validate:"required,gte=1,lte=15,dive,required"`
	}
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	conform.Strings(request)
	escape(request)

	if err := bind(request, context); err != nil {
		return request, err
	}

	if err := validate(request, context); err != nil {
		return request, err
	}

	return request, nil
}

func bind(request *Request, context echo.Context) error {
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

func validate(request *Request, context echo.Context) error {
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

func escape(request *Request) {
	sanitizer := bluemonday.StrictPolicy()

	request.Query = sanitizer.Sanitize(request.Query)
	request.OperationName = sanitizer.Sanitize(request.OperationName)

	for key, value := range request.Variables {
		switch actualValue := value.(type) {
		case int:
			break
		case float64:
			break
		case string:
			request.Variables[key] = sanitizer.Sanitize(actualValue)
		case bool:
			break
		default:
			break
		}
	}
}
