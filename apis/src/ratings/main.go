package main

import (
	"fmt"

	"ratings/controller"
	"ratings/parser"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	port := "3000"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &parser.RequestValidator{validator: validator.New()}

	e.POST("/ratings", controller.PostRatings)
	e.OPTIONS("/ratings", controller.OptionsRatings)
	e.OPTIONS("/", controller.OptionsRoot)

	fmt.Println("Started server on port", port)

	e.Server.Addr = ":" + port
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
