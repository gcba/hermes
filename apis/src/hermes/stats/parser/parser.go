package parser

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/microcosm-cc/bluemonday"
)

// Request holds the mapped fields from the request's JSON body
type (
	Request struct {
		Query     string                 `json:"query" validate:"required,gte=10,lte=5000" conform:"trim"`
		Variables map[string]interface{} `json:"variables" validate:"required"`
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
	request.Variables, _ = sanitizeMap(request.Variables, sanitizer).(map[string]interface{})
}

// From: https://gist.github.com/hvoecking/10772475, license: MIT
func sanitizeMap(obj interface{}, sanitizer *bluemonday.Policy) interface{} {
	original := reflect.ValueOf(obj)
	copy := reflect.New(original.Type()).Elem()

	sanitizeMapRecursive(copy, original, sanitizer)

	return copy.Interface()
}

// From: https://gist.github.com/hvoecking/10772475, license: MIT
func sanitizeMapRecursive(copy, original reflect.Value, sanitizer *bluemonday.Policy) {
	switch original.Kind() {
	case reflect.Interface:
		originalValue := original.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()

		sanitizeMapRecursive(copyValue, originalValue, sanitizer)
		copy.Set(copyValue)
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))

		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()

			sanitizeMapRecursive(copyValue, originalValue, sanitizer)
			copy.SetMapIndex(key, copyValue)
		}
	case reflect.String:
		originalValue := fmt.Sprintf("%v", original.Elem())
		sanitizedValue := strings.TrimSpace(sanitizer.Sanitize(originalValue))

		copy.SetString(sanitizedValue)
	default:
		copy.Set(original)
	}
}
