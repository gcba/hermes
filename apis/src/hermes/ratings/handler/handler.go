package handler

import (
	base "hermes/handler"
	"hermes/ratings/parser"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func Handler(port int, routes map[string]echo.HandlerFunc) *echo.Echo {
	e := base.Handler(port, "HERMES_RATINGS_ENV", "HERMES_RATINGS_PUBLICKEY")
	validate := validator.New()

	parser.RegisterCustomValidators(validate)
	e.OPTIONS("/", routes["OptionsRoot"])
	e.OPTIONS("/ratings", routes["OptionsRatings"])
	e.POST("/ratings", routes["PostRatings"])

	e.Validator = &RequestValidator{validator: validate}

	return e
}
