package controller

import (
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type databases struct {
	Read *gorm.DB,
	Write *gorm.DB
}

// PostRating saves a new rating to the database
func PostRating(context echo.Context) error {
	request := parser.Parse(context)
	dbs := &databases{Read: &GetReadDB(), Write: &GetWriteDB()}

	defer dbs.Read.Close()
	defer dbs.Write.Close()

	if rating, ok := newRating(request, dbs, context); !ok {
		err := "Rating creation failed"

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	return responses.PostResponse(http.StatusOK, context)
}

func OptionsRating(context echo.Context) error {
	endpoints := []responses.Endpoint{responses.Endpoints["PostRatings"]}

  	return responses.OptionsResponse(endpoints, context)
}

func hasAppUser(request *parser.Request) bool {
	appuser := request.User

	if appuser.Name == "" || appuser.Email == "" || appuser.MiBAID == "" {
		return false
	}

	return true
}

func addAppUser(request *parser.Request, rating *models.Rating, dbs *databases) {
	if hasAppUser(request) {
		appUser, ok := getAppUser(request, dbs); !ok {
			err := "Error trying to get an app user from the database"

			return responses.ErrorResponse(http.StatusInternalServerError, err, context)
		}

		rating.AppUserID = appUser.ID
	}
}

func hasDevice(request *parser.Request) bool {
	device := request.Device

	if device.Name == "" || device.Brand == "" {
		return false
	}

	return true
}

func addDevice(request *parser.Request, rating *models.Rating, dbs *databases) {
	brand, brandOk := getBrand(request, dbs); !brandOk {
		err := "Error trying to get a brand from the database"

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	device, deviceOk := getDevice(request, brand, platform, dbs); !deviceOk {
		err := "Error trying to get a device from the database"

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	rating.DeviceID = device.ID
}

func hasBrowser(request *parser.Request) bool {
	browser := request.Browser

	if browser.Name == "" {
		return false
	}

	return true
}

func addBrowser(request *parser.Request, rating *models.Rating, dbs *databases) {
	browser, ok := getBrowser(request, dbs); !ok {
		err := "Error trying to get a browser from the database"

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	rating.BrowserID = browser.ID
}

func getApp(request *parser.Request, db *gorm.DB, context echo.Context) models.App {
	app, appErr := models.GetApp(request.App.Key, db)

	if appErr != nil {
		err := fmt.Sprintf("Error getting app: %s", appErr)

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	return app
}

func getPlatform(request *parser.Request, db *gorm.DB, context echo.Context) models.Platform {
	platform, platformErr := models.GetPlatform(request.Platform.Key, db)

	if platformErr != nil {
		err := fmt.Sprintf("Error getting platform: %s", platform)

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	return platform
}

func getRange(request *parser.Request, db *gorm.DB, context echo.Context) models.Range {
	rangeRecord, rangeErr := models.GetRange(request.Range, db)

	if rangeErr != nil {
		err := fmt.Sprintf("Error getting range: %s", rangeErr)

		return responses.ErrorResponse(http.StatusInternalServerError, err, context)
	}

	return rangeRecord
}

func getBrowser(request *parser.Request, dbs *databases) (models.Browser, bool) {
	getResult := models.GetBrowser(request.Browser.Name, dbs.Read)

	if getResult.RecordNotFound() {
		browser := Browser{Name: request.Browser.Name}
		createResult := models.CreateBrowser(&browser, dbs.Write)

		if len(createResult.GetErrors()) > 0 {
			// TODO: Handle errors

			for value, err := range getResult.GetErrors() {
				fmt.Println("Error creating a Browser:", value)
			}
		}

		fmt.Println("Created a new Browser:", createResult.Value)

		return createResult.Value, true
	}
	else if len(getResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error getting a Browser:", value)
		}

		return Browser{}, false
	}

	return getResult.Value, true
}

func getBrand(request *parser.Request, dbs *databases) (models.Brand, bool) {
	getResult := models.GetBrand(request.Browser.Name, dbs.Read)

	if getResult.RecordNotFound() {
		brand := Brand{Name: request.Brand}
		createResult := models.CreateBrand(&brand, dbs.Write)

		if len(createResult.GetErrors()) > 0 {
			// TODO: Handle errors

			for value, err := range getResult.GetErrors() {
				fmt.Println("Error creating a Brand:", value)
			}
		}

		fmt.Println("Created a new Brand:", createResult.Value)

		return createResult.Value, true
	}
	else if len(getResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error getting a Brand:", value)
		}

		return Brand{}, false
	}

	return getResult.Value, true
}

func getDevice(request *parser.Request, brand *models.Brand, platform *models.Platform, dbs *databases) (models.Device, bool) {
	getResult := models.GetDevice(request.Device.Name, request.Brand, dbs.Read)

	if getResult.RecordNotFound() {
		device := Device{
				Name: request.Device.Name,
				ScreenWidth: request.Device.Screen.Width,
				ScreenHeight: request.Device.Screen.Height,
				PPI: request.Device.Screen.PPI,
				BrandID: brand.ID,
				PlatformID: platform.ID
			}

		createResult := models.CreateDevice(&device, dbs.Write)

		if len(createResult.GetErrors()) > 0 {
			// TODO: Handle errors

			for value, err := range getResult.GetErrors() {
				fmt.Println("Error creating a Device:", value)
			}
		}

		fmt.Println("Created a new Device:", createResult.Value)

		return createResult.Value, true
	}
	else if len(getResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error getting a Device:", value)
		}

		return Device{}, false
	}

	return getResult.Value, true
}

func getAppUser(request *parser.Request, dbs *databases) (models.AppUser, bool) {
	getResult := models.GetAppUser(request.User.MiBAID, dbs.Read)

	if getResult.RecordNotFound() {
		appuser := AppUser{
			Name: request.User.Name,
			Email: request.User.Email,
			MiBAID: request.User.MiBAID
		}

		createResult := models.CreateAppUser(&appuser, dbs.Write)

		if len(createResult.GetErrors()) > 0 {
			// TODO: Handle errors

			for value, err := range getResult.GetErrors() {
				fmt.Println("Error creating an AppUser:", value)
			}
		}

		fmt.Println("Created a new AppUser:", createResult.Value)

		return createResult.Value, true
	}
	else if len(getResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error getting an AppUser:", value)
		}

		return AppUser{}, false
	}

	return getResult.Value, true
}

func newRating(request *parser.Request, dbs *databases, context echo.Context) (models.Rating, bool) {
	app := getApp(request, dbs.Read, context)
	platform := getPlatform(request, dbs.Read, context)
	rangeRecord := getRange(request, dbs.Read, context)
	hasMessage := false

	if len(request.Comment) > 0 {
		hasMessage = true
	}

	rating := &Rating{
		Rating: request.Rating,
		Description: request.Description,
		AppVersion: request.App.Version,
		PlatformVersion: request.Platform.Version,
		BrowserVersion: request.Browser.Version,
		HasMessage: hasMessage,
		AppID: app.ID,
		RangeID: rangeRecord.ID,
		PlatformID: platform.ID
	}

	if hasAppUser(request) {
		addAppUser(request, rating, dbs)
	}

	if hasDevice(request) {
		addDevice(request, rating, dbs)
	}

	if hasBrowser(request) {
		addBrowser(request, rating, dbs)
	}

	createResult := models.CreateRating(rating, dbs.Write)

	if len(createResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error creating a Rating:", value)
		}
	}

	fmt.Println("Created a new Rating:", createResult.Value)

	return createResult.Value, true
}