package handler

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"hermes/stats/responses"

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

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		if context.Request().Method == echo.OPTIONS && hasContentTypeHeader(context) {
			message = "OPTIONS requests must have no body"

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		if context.Request().Method == echo.POST && !hasContentTypeHeader(context) {
			message = "Content-Type header is missing"

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		return next(context)
	}
}

func notAcceptableMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var message string

		if !isValidAcceptHeader(context) {
			message = "JSON responses must be accepted"

			return echo.NewHTTPError(http.StatusNotAcceptable, []string{message})
		}

		return next(context)
	}
}

func unsupportedMediaTypeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var message string

		if context.Request().Method == echo.POST {
			if !isValidContentTypeHeader(context) || !isValidCharacterEncoding(context) {
				message = "Request body must be UTF-8 encoded JSON"

				return echo.NewHTTPError(http.StatusUnsupportedMediaType, []string{message})
			}
		}

		return next(context)
	}
}

func notImplementedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if context.Request().Method != echo.POST && context.Request().Method != echo.OPTIONS {
			return echo.NewHTTPError(http.StatusNotImplemented)
		}

		return next(context)
	}
}

func corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		context.Response().Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		context.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

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
	accepted := "application/json"

	if header := context.Request().Header.Get("Accept"); strings.Contains(strings.ToLower(header), accepted) || header == "*/*" {
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
	text := "text/plain"
	json := "application/json"

	if header := context.Request().Header.Get("Content-Type"); strings.Contains(strings.ToLower(header), text) || strings.Contains(strings.ToLower(header), json) {
		return true
	}

	return false
}

func isValidCharacterEncoding(context echo.Context) bool {
	charset := "charset=utf-8"

	if header := context.Request().Header.Get("Content-Type"); strings.HasSuffix(strings.ToLower(header), charset) {
		return true
	}

	return false
}

func Handler(port int, handlers map[string]echo.HandlerFunc) http.Handler {
	e := echo.New()
	validate := validator.New()
	env := os.Getenv("HERMES_STATS_ENV")

	if env == "DEV" {
		e.Logger.SetLevel(log.DEBUG)
		e.Debug = true
	} else {
		e.Pre(middleware.HTTPSRedirect())
		e.Logger.SetLevel(log.ERROR)
	}

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("20K"))
	e.Use(notImplementedMiddleware)
	e.Use(notAcceptableMiddleware)
	e.Use(badRequestMiddleware)
	e.Use(unsupportedMediaTypeMiddleware)
	e.Use(corsMiddleware)

	e.POST("/stats", handlers["PostStats"])

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validate}
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}
