package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateRating(c echo.Context) error {
	r := new(Request)

	if err := c.Bind(r); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}
