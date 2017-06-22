package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func Handler(port int, handlers map[string]echo.HandlerFunc) http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("10K"))

	e.Server.Addr = ":" + strconv.Itoa(port)
	e.Validator = &RequestValidator{validator: validator.New()} // TODO: Move this to parser.AttachValidator()

	e.OPTIONS("/", handlers["OptionsRoot"])
	e.OPTIONS("/ratings", handlers["OptionsRatings"])
	e.POST("/ratings", handlers["PostRatings"])

	return e
}
