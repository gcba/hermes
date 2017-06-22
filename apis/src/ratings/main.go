package main

import (
	"fmt"
	"strconv"

	"ratings/controller"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func Handler(port int) http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("10K"))

	e.Server.Addr = ":" + strconv.Itoa(port)
	e.Validator = &RequestValidator{validator: validator.New()} // TODO: Move this to parser.AttachValidator()

	e.OPTIONS("/", controller.OptionsRoot)
	e.OPTIONS("/ratings", controller.OptionsRatings)
	e.POST("/ratings", controller.PostRatings)

	return e
}

func main() {
	port := 3000
	handler := Handler(port).(*echo.Echo) // Casting via type assertion

	fmt.Println("Started server on port", strconv.Itoa(port))

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}
