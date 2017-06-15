package controller

import (
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"

	"github.com/labstack/echo"
)

func Create(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	readDb := GetReadDB()
	defer readDb.Close()

	brand, err := models.GetBrand("Google", readDb)

	fmt.Println(request.Rating)

	return context.JSON(http.StatusOK, brand)
}
