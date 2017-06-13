package main

import (
	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Request struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ratings", createRating)

	e.Server.Addr = ":3000"
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func createRating(c echo.Context) error {

}
