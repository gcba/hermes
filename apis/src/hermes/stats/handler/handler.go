package handler

import (
	"os"
	"strconv"

	"hermes/middlewares"
	"hermes/responses"
	"hermes/stats/controller"
	"hermes/stats/schema"

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

func Handler(port int) *echo.Echo {
	schema.Parse()

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
	e.Use(middlewares.NotImplementedMiddleware)
	e.Use(middlewares.NotAcceptableMiddleware)
	e.Use(middlewares.BadRequestMiddleware)
	e.Use(middlewares.UnsupportedMediaTypeMiddleware)
	e.Use(middlewares.CorsMiddleware)

	e.POST("/stats", controller.PostStats)

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validate}
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}
