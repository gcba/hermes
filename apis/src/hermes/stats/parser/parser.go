package parser

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	base "hermes/parser"

	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
)

// Request holds the mapped fields from the request's JSON body
type (
	Request struct {
		Query     string                 `json:"query" validate:"required,gte=10,lte=5000" conform:"trim"`
		Variables map[string]interface{} `json:"variables" validate:"required"`
	}
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(context echo.Context) (*Request, error) {
	rawRequest, err := base.Parse(newRequest, escape, context)
	request, castOk := rawRequest.(*Request)

	if err != nil {
		return nil, err
	}

	if !castOk {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return request, nil
}

func newRequest() interface{} {
	return new(Request)
}

func escape(request interface{}) error {
	sanitizer := bluemonday.StrictPolicy()

	if data, ok := request.(*Request); ok {
		data.Query = sanitizer.Sanitize(data.Query)
		data.Variables, _ = sanitizeMap(data.Variables, sanitizer).(map[string]interface{})

		return nil
	}

	return echo.NewHTTPError(http.StatusInternalServerError)
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
