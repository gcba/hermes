package main

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/labstack/echo"
)

func main() {
	// os.Getenv("S3_BUCKET")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
