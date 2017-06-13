package main

import (
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Request struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	User    struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Platform struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Device struct {
		Name   string `json:"name"`
		Brand  string `json:"brand"`
		Screen struct {
			Width  int     `json:"width"`
			Height int     `json:"height"`
			PPI    float32 `json:"ppi"`
		}
	}
	Browser struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
}

func main() {
	e := echo.New()
	port := "3000"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ratings", createRating)

	fmt.Println("Started server on port", port)

	e.Server.Addr = ":" + port
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func createRating(c echo.Context) error {
	r := new(Request)

	if err := c.Bind(r); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}
