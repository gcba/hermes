package controller

import (
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Create saves a new rating to the database
func Create(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	readDb := GetReadDB()
	defer readDb.Close()
	writeDb := GetWriteDB()
	defer writeDb.Close()

	appRecord, rangeRecord, platformRecord := getBaseFields(request, readDb)

	fmt.Println("appRecord:", appRecord)
	fmt.Println("rangeRecord:", rangeRecord)

	return context.JSON(http.StatusOK, platformRecord)
}

func getBaseFields(request *parser.Request, db *gorm.DB) (models.App, models.Range, models.Platform) {
	appRecord, appErr := models.GetApp(request.App, db)

	if appErr != nil {
		// TODO: Dispatch error response
		fmt.Println("Error getting app:", appErr)
	}

	rangeRecord, rangeErr := models.GetRange(request.Range, db)

	if rangeErr != nil {
		// TODO: Dispatch error response
		fmt.Println("Error getting range:", rangeErr)
	}

	platformRecord, platformErr := models.GetPlatform(request.Platform.Key, db)

	if platformErr != nil {
		// TODO: Dispatch error response
		fmt.Println("Error getting platform:", platformErr)
	}

	return appRecord, rangeRecord, platformRecord
}

func hasAppUser(request *parser.Request) bool {
	appuser := request.User

	if appuser.Name == "" || appuser.Email == "" || appuser.MiBAID == "" {
		return false
	}

	return true
}

func hasDevice(request *parser.Request) bool {
	device := request.Device

	if device.Name == "" || device.Brand == "" {
		return false
	}

	return true
}

func hasBrowser(request *parser.Request) bool {
	browser := request.Browser

	if browser.Name == "" {
		return false
	}

	return true
}
