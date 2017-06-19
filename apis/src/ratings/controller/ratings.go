package controller

import (
	"errors"
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"
	"ratings/responses"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type databases struct {
	Read  *gorm.DB
	Write *gorm.DB
}

// PostRatings saves a new rating to the database
func PostRatings(context echo.Context) error {
	request := parser.Parse(context)
	readDB := GetReadDB()
	writeDB := GetWriteDB()
	dbs := &databases{Read: readDB, Write: writeDB}

	defer readDB.Close()
	defer writeDB.Close()

	if err := newRating(request, dbs, context); err != nil {
		return err
	}

	return responses.PostResponse(http.StatusOK, context)
}

// OptionsRatings returns the response of the `OPTIONS /ratings` endpoint
func OptionsRatings(context echo.Context) error {
	endpoints := []responses.Endpoint{responses.Endpoints["PostRatings"]}

	return responses.OptionsResponse(endpoints, context)
}

func handleError(err error, context echo.Context) {
	responses.ErrorResponse(http.StatusInternalServerError, err.Error(), context)
}

func handleErrors(errors []error, context echo.Context) {
	stringErrors := make([]string, len(errors))

	for index, err := range errors {
		stringErrors[index] = err.Error()
	}

	responses.ErrorsResponse(http.StatusInternalServerError, stringErrors, context)
}

/*
*
* App
*
 */

func getApp(request *parser.Request, db *gorm.DB, context echo.Context) models.App {
	result := models.GetApp(request.App.Key, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		handleErrors(errorList, context)
	} else if value, ok := result.Value.(models.App); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get an app from the database"), context)
	}

	return models.App{}
}

/*
*
* Platform
*
 */

func getPlatform(request *parser.Request, db *gorm.DB, context echo.Context) models.Platform {
	result := models.GetPlatform(request.Platform.Key, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		handleErrors(errorList, context)
	} else if value, ok := result.Value.(models.Platform); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get a platform from the database"), context)
	}

	return models.Platform{}
}

/*
*
* Range
*
 */

func getRange(request *parser.Request, db *gorm.DB, context echo.Context) models.Range {
	result := models.GetRange(request.Range, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		handleErrors(errorList, context)
	} else if value, ok := result.Value.(models.Range); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get a range from the database"), context)
	}

	return models.Range{}
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

func getAppUser(request *parser.Request, dbs *databases, context echo.Context) models.AppUser {
	getResult := models.GetAppUser(request.User.MiBAID, dbs.Read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		appuser := &models.AppUser{
			Name:   request.User.Name,
			Email:  request.User.Email,
			MiBAID: request.User.MiBAID}

		createResult := models.CreateAppUser(appuser, dbs.Write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			handleErrors(createErrorList, context)
		} else if value, ok := createResult.Value.(models.AppUser); ok {
			fmt.Println("Created a new AppUser:", createResult.Value)

			return value
		} else {
			handleError(errors.New("Error trying to create an app user"), context)
		}

		return models.AppUser{}
	}

	if len(getErrorList) > 0 {
		handleErrors(getErrorList, context)
	} else if value, ok := getResult.Value.(models.AppUser); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get an app user from the database"), context)
	}

	return models.AppUser{}
}

func attachAppUser(request *parser.Request, rating *models.Rating, dbs *databases, context echo.Context) {
	if hasAppUser(request) {
		appUser := getAppUser(request, dbs, context)
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

func getBrowser(request *parser.Request, dbs *databases, context echo.Context) models.Browser {
	getResult := models.GetBrowser(request.Browser.Name, dbs.Read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &models.Browser{Name: request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.Write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			handleErrors(createErrorList, context)
		} else if value, ok := createResult.Value.(models.Browser); ok {
			fmt.Println("Created a new Browser:", createResult.Value)

			return value
		} else {
			handleError(errors.New("Error trying to create a browser"), context)
		}

		return models.Browser{}
	}

	if len(getErrorList) > 0 {
		handleErrors(getErrorList, context)
	} else if value, ok := getResult.Value.(models.Browser); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get a browser from the database"), context)
	}

	return models.Browser{}
}

func attachBrowser(request *parser.Request, rating *models.Rating, dbs *databases, context echo.Context) {
	browser := getBrowser(request, dbs, context)
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

func getDevice(request *parser.Request, brand *models.Brand, platform *models.Platform, dbs *databases, context echo.Context) models.Device {
	getResult := models.GetDevice(request.Device.Name, brand.ID, dbs.Read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		device := &models.Device{
			Name:         request.Device.Name,
			ScreenWidth:  request.Device.Screen.Width,
			ScreenHeight: request.Device.Screen.Height,
			PPI:          request.Device.Screen.PPI,
			BrandID:      brand.ID,
			PlatformID:   platform.ID}

		createResult := models.CreateDevice(device, dbs.Write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			handleErrors(createErrorList, context)
		} else if value, ok := createResult.Value.(models.Device); ok {
			fmt.Println("Created a new Device:", createResult.Value)

			return value
		} else {
			handleError(errors.New("Error trying to create a device"), context)
		}

		return models.Device{}
	}

	if len(getErrorList) > 0 {
		handleErrors(getErrorList, context)
	} else if value, ok := getResult.Value.(models.Device); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get a device from the database"), context)
	}

	return models.Device{}
}

func attachDevice(request *parser.Request, rating *models.Rating, platform *models.Platform, dbs *databases, context echo.Context) {
	brand := getBrand(request, dbs, context)
	device := getDevice(request, &brand, platform, dbs, context)
	rating.DeviceID = device.ID
}

/*
*
* Brand
*
 */

func getBrand(request *parser.Request, dbs *databases, context echo.Context) models.Brand {
	getResult := models.GetBrand(request.Browser.Name, dbs.Read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &models.Brand{Name: request.Device.Brand}
		createResult := models.CreateBrand(brand, dbs.Write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			handleErrors(createErrorList, context)
		} else if value, ok := createResult.Value.(models.Brand); ok {
			fmt.Println("Created a new Brand:", createResult.Value)

			return value
		} else {
			handleError(errors.New("Error trying to create a brand"), context)
		}

		return models.Brand{}
	}

	if len(getErrorList) > 0 {
		handleErrors(getErrorList, context)
	} else if value, ok := getResult.Value.(models.Brand); ok {
		return value
	} else {
		handleError(errors.New("Error trying to get a brand from the database"), context)
	}

	return models.Brand{}
}

/*
*
* Rating
*
 */

func newRating(request *parser.Request, dbs *databases, context echo.Context) error {
	app := getApp(request, dbs.Read, context)
	platform := getPlatform(request, dbs.Read, context)
	rangeRecord := getRange(request, dbs.Read, context)
	hasMessage := false

	if len(request.Comment) > 0 {
		hasMessage = true
	}

	rating := &models.Rating{
		Rating:          request.Rating,
		Description:     request.Description,
		AppVersion:      request.App.Version,
		PlatformVersion: request.Platform.Version,
		BrowserVersion:  request.Browser.Version,
		HasMessage:      hasMessage,
		AppID:           app.ID,
		RangeID:         rangeRecord.ID,
		PlatformID:      platform.ID}

	if hasAppUser(request) {
		attachAppUser(request, rating, dbs, context)
	}

	if hasDevice(request) {
		attachDevice(request, rating, &platform, dbs, context)
	}

	if hasBrowser(request) {
		attachBrowser(request, rating, dbs, context)
	}

	result := models.CreateRating(rating, dbs.Write)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		handleErrors(errorList, context)
	} else if value, ok := result.Value.(models.Rating); ok {
		fmt.Println("Created a new Rating:", value)

		return nil
	} else {
		handleError(errors.New("Error trying to create a rating"), context)
	}

	return errors.New("Could not create a new rating")
}
