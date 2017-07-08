package handler

import (
	"net/http"
	"strconv"
	"strings"

	"ratings/responses"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func badRequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var message string

		if !hasAcceptHeader(context) {
			message = "Accept header is missing"

			return responses.ErrorResponse(http.StatusBadRequest, message, context)
		}

		if !isValidAcceptHeader(context) {
			message = "Accept header must equal 'application/json'"

			return responses.ErrorResponse(http.StatusBadRequest, message, context)
		}

		if context.Request().Method == echo.POST && !hasContentTypeHeader(context) {
			message = "Content-Type header is missing"

			return responses.ErrorResponse(http.StatusBadRequest, message, context)
		}

		if context.Request().Method == echo.POST && !isValidContentTypeHeader(context) {
			message = "Content-Type header must equal 'application/json; charset=UTF-8'"

			return responses.ErrorResponse(http.StatusBadRequest, message, context)
		}

		return next(context)
	}
}

func notImplementedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if context.Request().Method != echo.POST && context.Request().Method != echo.OPTIONS {
			return responses.ErrorResponse(http.StatusNotImplemented, "", context)
		}

		return next(context)
	}
}

func hasAcceptHeader(context echo.Context) bool {
	if header := context.Request().Header.Get("Accept"); strings.TrimSpace(header) != "" {
		return true
	}

	return false
}

func isValidAcceptHeader(context echo.Context) bool {
	if header := context.Request().Header.Get("Accept"); header == "application/json" || header == "*/*" {
		return true
	}

	return false
}

func hasContentTypeHeader(context echo.Context) bool {
	if header := context.Request().Header.Get("Content-Type"); strings.TrimSpace(header) != "" {
		return true
	}

	return false
}

func isValidContentTypeHeader(context echo.Context) bool {
	validHeader := "application/json; charset=utf-8"

	if header := context.Request().Header.Get("Content-Type"); strings.ToLower(header) == validHeader {
		return true
	}

	return false
}

func Handler(port int, handlers map[string]echo.HandlerFunc) http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(notImplementedMiddleware)
	e.Use(badRequestMiddleware)

	e.OPTIONS("/", handlers["OptionsRoot"])
	e.OPTIONS("/ratings", handlers["OptionsRatings"])
	e.POST("/ratings", handlers["PostRatings"])

	e.Logger.SetLevel(log.ERROR)

	e.Server.Addr = ":" + strconv.Itoa(port)
	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validator.New()}

	return e
}
