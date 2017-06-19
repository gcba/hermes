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
func PostRatings(context echo.Context) error {
	request := parser.Parse(context)
	dbs := &databases{Read: &GetReadDB(), Write: &GetWriteDB()}

	defer dbs.Read.Close()
	defer dbs.Write.Close()

	rating := newRating(request, dbs, context)

	return responses.PostResponse(http.StatusOK, context)
}

// OptionsRating returns OPTIONS /ratings response
func OptionsRatings(context echo.Context) error {
	endpoints := []responses.Endpoint{responses.Endpoints["PostRatings"]}

  	return responses.OptionsResponse(endpoints, context)
}

func handleErrors(errors []error, context) {
	stringErrors := []string

	for value, err := range errors {
		append(stringErrors, value.Error())
	}

	return responses.ErrorsResponse(http.StatusInternalServerError, stringErrors, context)
}

/*
*
* App
*
*/

func getApp(request *parser.Request, db *gorm.DB, context echo.Context) models.App {
	result := models.GetApp(request.App.Key, db)
	errors := result.GetErrors()

	if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return result
}

/*
*
* Platform
*
*/

func getPlatform(request *parser.Request, db *gorm.DB, context echo.Context) models.Platform {
	result := models.GetPlatform(request.Platform.Key, db)
	errors := result.GetErrors()

	if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return result
}

/*
*
* Range
*
*/

func getRange(request *parser.Request, db *gorm.DB, context echo.Context) models.Range {
	result := models.GetRange(request.Range, db)
	errors := result.GetErrors()

	if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return result
}

/*
*
* AppUser
*
*/

func hasAppUser(request *parser.Request) bool {
	appuser := request.User

	if appuser.Name == "" || appuser.Email == "" || appuser.MiBAID == "" {
		return false
	}

	return true
}

func getAppUser(request *parser.Request, dbs *databases) models.AppUser {
	getResult := models.GetAppUser(request.User.MiBAID, dbs.Read)
	errors := getResult.GetErrors()

	if getResult.RecordNotFound() {
		appuser := &AppUser{
			Name: request.User.Name,
			Email: request.User.Email,
			MiBAID: request.User.MiBAID
		}

		createResult := models.CreateAppUser(appuser, dbs.Write)
		errors := createResult.GetErrors()

		if len(errors) > 0 {
			handleErrors(errors, context)
		}

		fmt.Println("Created a new AppUser:", createResult.Value)

		return createResult.Value,
	}
	else if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return getResult.Value
}

func attachAppUser(request *parser.Request, rating *models.Rating, dbs *databases) {
	if hasAppUser(request) {
		appUser:= getAppUser(request, dbs)
		rating.AppUserID = appUser.ID
	}
}

/*
*
* Browser
*
*/

func hasBrowser(request *parser.Request) bool {
	browser := request.Browser

	if browser.Name == "" {
		return false
	}

	return true
}

func getBrowser(request *parser.Request, dbs *databases) models.Browser {
	getResult := models.GetBrowser(request.Browser.Name, dbs.Read)
	errors := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &Browser{Name: request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.Write)
		errors := createResult.GetErrors()

		if len(errors) > 0 {
			handleErrors(errors, context)
		}

		fmt.Println("Created a new Browser:", createResult.Value)

		return createResult.Value
	}
	else if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return getResult.Value
}

func attachBrowser(request *parser.Request, rating *models.Rating, dbs *databases) {
	browser := getBrowser(request, dbs)
	rating.BrowserID = browser.ID
}

/*
*
* Device
*
*/

func hasDevice(request *parser.Request) bool {
	device := request.Device

	if device.Name == "" || device.Brand == "" {
		return false
	}

	return true
}

func getDevice(request *parser.Request, brand *models.Brand, platform *models.Platform, dbs *databases) models.Device {
	getResult := models.GetDevice(request.Device.Name, request.Brand, dbs.Read)
	errors := getResult.GetErrors()

	if getResult.RecordNotFound() {
		device := &Device{
				Name: request.Device.Name,
				ScreenWidth: request.Device.Screen.Width,
				ScreenHeight: request.Device.Screen.Height,
				PPI: request.Device.Screen.PPI,
				BrandID: brand.ID,
				PlatformID: platform.ID
			}

		createResult := models.CreateDevice(device, dbs.Write)
		errors := createResult.GetErrors()

		if len(errors) > 0 {
			handleErrors(errors, context)
		}

		fmt.Println("Created a new Device:", createResult.Value)

		return createResult.Value
	}
	else if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return getResult.Value
}

func attachDevice(request *parser.Request, rating *models.Rating, dbs *databases) {
	brand := getBrand(request, dbs)
	device := getDevice(request, brand, platform, dbs)
	rating.DeviceID = device.ID
}

/*
*
* Brand
*
*/

func getBrand(request *parser.Request, dbs *databases) models.Brand {
	getResult := models.GetBrand(request.Browser.Name, dbs.Read)
	errors := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &Brand{Name: request.Brand}
		createResult := models.CreateBrand(brand, dbs.Write)
		errors := createResult.GetErrors()

		if len(errors) > 0 {
			handleErrors(errors, context)
		}

		fmt.Println("Created a new Brand:", createResult.Value)

		return createResult.Value
	}
	else if len(errors) > 0 {
		handleErrors(errors, context)
	}

	return getResult.Value
}

/*
*
* Rating
*
*/

func newRating(request *parser.Request, dbs *databases, context echo.Context) *models.Rating {
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
		attachAppUser(request, rating, dbs)
	}

	if hasDevice(request) {
		attachDevice(request, rating, dbs)
	}

	if hasBrowser(request) {
		attachBrowser(request, rating, dbs)
	}

	result := models.CreateRating(rating, dbs.Write)
	errors := result.GetErrors()

	else if len(errors) > 0 {
		handleErrors(errors, context)
	}

	fmt.Println("Created a new Rating:", result.Value)

	return result.Value
}