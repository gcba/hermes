package main

import (
	"fmt"

	"ratings/controller"

	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	port := "3000"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ratings", controller.PostRating)
	e.OPTIONS("/ratings", controller.OptionsRating)
	e.OPTIONS("/", controller.OptionsRoot)

	fmt.Println("Started server on port", port)

	e.Server.Addr = ":" + port
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
