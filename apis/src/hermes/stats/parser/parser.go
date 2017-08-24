package parser

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/microcosm-cc/bluemonday"
)

// Request holds the mapped fields from the request's JSON body

type (
	operation struct {
		Condition string `json:"condition" validate:"required,gte=1,lte=3,alpha" conform:"trim,upper"` // TODO: Add custom validator
		Field     *Field `json:"field" validate:"omitempty"`
	}

	variables struct {
		Field Field `json:"field" validate:"required"`
	}

	Field struct {
		Name     string      `json:"name" validate:"required,gte=3,lte=30,contains=.,excludesall= " conform:"trim,lower"` // TODO: Add custom validator
		Operator string      `json:"operator" validate:"omitempty,gte=1,lte=3,alpha" conform:"trim,upper"`                // TODO: Add custom validator
		Value    interface{} `json:"value" validate:"required"`
		Next     *operation  `json:"next" validate:"omitempty"`
	}

	Request struct {
		Query     string    `json:"query" validate:"required,gte=10,lte=5000" conform:"trim"`
		Variables variables `json:"variables" validate:"required"`
		// Variables map[string]interface{} `json:"variables" validate:"required"`
	}
)

// TODO: Consider extracting the common parts into its own package

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
	request.Variables.Field.Name = sanitizer.Sanitize(request.Variables.Field.Name)

	switch value := request.Variables.Field.Value.(type) {
	case string:
		request.Variables.Field.Value = sanitizer.Sanitize(value)
	default:
		break
	}
}
