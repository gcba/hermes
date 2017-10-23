package handler

import (
	base "hermes/handler"
	"hermes/stats/schema"

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
	schema.Parse()

	e := base.Handler(port, "HERMES_STATS_ENV", "HERMES_STATS_PUBLICKEY")

	e.POST("/stats", routes["PostStats"])

	e.Validator = &RequestValidator{validator: validator.New()}

	return e
}
