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

type (
	frame struct {
		request *parser.Request
		context echo.Context
	}

	databases struct {
		read  *gorm.DB
		write *gorm.DB
	}
)

func errorResponse(err error, context echo.Context) error {
	return responses.ErrorResponse(http.StatusInternalServerError, err.Error(), context)
}

func errorsResponse(errors []error, context echo.Context) error {
	stringErrors := make([]string, len(errors))

	for index, err := range errors {
		stringErrors[index] = err.Error()
	}

	return responses.ErrorsResponse(http.StatusInternalServerError, stringErrors, context)
}

/*
*
* App
*
 */
func getApp(db *gorm.DB, frame *frame) (models.App, error) {
	result := models.GetApp(frame.request.App.Key, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		return models.App{}, errorsResponse(errorList, frame.context)
	} else if value, ok := result.Value.(models.App); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get an app from the database"), frame.context)

		return models.App{}, err
	}
}

/*
*
* Platform
*
 */
func getPlatform(db *gorm.DB, frame *frame) (models.Platform, error) {
	result := models.GetPlatform(frame.request.Platform.Key, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		return models.Platform{}, errorsResponse(errorList, frame.context)
	} else if value, ok := result.Value.(models.Platform); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get a platform from the database"), frame.context)

		return models.Platform{}, err
	}
}

/*
*
* Range
*
 */
func getRange(db *gorm.DB, frame *frame) (models.Range, error) {
	result := models.GetRange(frame.request.Range, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		return models.Range{}, errorsResponse(errorList, frame.context)
	} else if value, ok := result.Value.(models.Range); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get a range from the database"), frame.context)

		return models.Range{}, err
	}
}

/*
*
* AppUser
*
 */
func hasAppUser(request *parser.Request) bool {
	appuser := request.User
	nameLength := len(appuser.Name)
	emailLength := len(appuser.Email)
	mibaIDLength := len(appuser.MiBAID)

	if nameLength == 0 || (emailLength == 0 && mibaIDLength == 0) {
		return false
	}

	return true
}

func getAppUser(dbs *databases, frame *frame) (models.AppUser, error) {
	var getResult *gorm.DB

	if hasMibaID := len(frame.request.User.MiBAID); hasMibaID > 0 {
		getResult = models.GetAppUser(frame.request.User.MiBAID, dbs.read)
	} else {
		getResult = models.GetAppUserByEmail(frame.request.User.Email, dbs.read)
	}

	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		appuser := &models.AppUser{
			Name:   frame.request.User.Name,
			Email:  frame.request.User.Email,
			MiBAID: frame.request.User.MiBAID}

		createResult := models.CreateAppUser(appuser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			return models.AppUser{}, errorsResponse(createErrorList, frame.context)
		} else if value, ok := createResult.Value.(models.AppUser); ok {
			fmt.Println("Created a new AppUser:", createResult.Value)

			return value, nil
		} else {
			err := errorResponse(errors.New("Error trying to create an app user"), frame.context)

			return models.AppUser{}, err
		}
	}

	if len(getErrorList) > 0 {
		return models.AppUser{}, errorsResponse(getErrorList, frame.context)
	} else if value, ok := getResult.Value.(models.AppUser); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get an app user from the database"), frame.context)

		return models.AppUser{}, err
	}
}

func attachAppUser(rating *models.Rating, dbs *databases, frame *frame) error {
	appUser, err := getAppUser(dbs, frame)

	if err != nil {
		rating.AppUserID = appUser.ID
	}

	return err
}

/*
*
* Browser
*
 */
func hasBrowser(request *parser.Request) bool {
	browser := request.Browser

	if len(browser.Name) == 0 {
		return false
	}

	return true
}

func getBrowser(dbs *databases, frame *frame) (models.Browser, error) {
	getResult := models.GetBrowser(frame.request.Browser.Name, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &models.Browser{Name: frame.request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			return models.Browser{}, errorsResponse(createErrorList, frame.context)
		} else if value, ok := createResult.Value.(models.Browser); ok {
			fmt.Println("Created a new Browser:", createResult.Value)

			return value, nil
		} else {
			err := errorResponse(errors.New("Error trying to create a browser"), frame.context)

			return models.Browser{}, err
		}
	}

	if len(getErrorList) > 0 {
		return models.Browser{}, errorsResponse(getErrorList, frame.context)
	} else if value, ok := getResult.Value.(models.Browser); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get a browser from the database"), frame.context)

		return models.Browser{}, err
	}
}

func attachBrowser(rating *models.Rating, dbs *databases, frame *frame) error {
	browser, err := getBrowser(dbs, frame)

	if err != nil {
		rating.BrowserID = browser.ID
	}

	return err
}

/*
*
* Device
*
 */
func hasDevice(request *parser.Request) bool {
	device := request.Device
	nameLength := len(device.Name)
	brandLength := len(device.Brand)

	if nameLength == 0 && brandLength == 0 {
		return false
	}

	return true
}

func getDevice(brand *models.Brand, platform *models.Platform, dbs *databases, frame *frame) (models.Device, error) {
	getResult := models.GetDevice(frame.request.Device.Name, brand.ID, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		device := &models.Device{
			Name:         frame.request.Device.Name,
			ScreenWidth:  frame.request.Device.Screen.Width,
			ScreenHeight: frame.request.Device.Screen.Height,
			PPI:          frame.request.Device.Screen.PPI,
			BrandID:      brand.ID,
			PlatformID:   platform.ID}

		createResult := models.CreateDevice(device, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			return models.Device{}, errorsResponse(createErrorList, frame.context)
		} else if value, ok := createResult.Value.(models.Device); ok {
			fmt.Println("Created a new Device:", createResult.Value)

			return value, nil
		} else {
			err := errorResponse(errors.New("Error trying to create a device"), frame.context)

			return models.Device{}, err
		}
	}

	if len(getErrorList) > 0 {
		return models.Device{}, errorsResponse(getErrorList, frame.context)
	} else if value, ok := getResult.Value.(models.Device); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get a device from the database"), frame.context)

		return models.Device{}, err
	}
}

func attachDevice(rating *models.Rating, platform *models.Platform, dbs *databases, frame *frame) error {
	brand, brandErr := getBrand(dbs, frame)

	if brandErr != nil {
		device, deviceErr := getDevice(&brand, platform, dbs, frame)

		if deviceErr != nil {
			rating.DeviceID = device.ID
		}

		return deviceErr
	}

	return brandErr
}

/*
*
* Brand
*
 */
func getBrand(dbs *databases, frame *frame) (models.Brand, error) {
	getResult := models.GetBrand(frame.request.Browser.Name, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &models.Brand{Name: frame.request.Device.Brand}
		createResult := models.CreateBrand(brand, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 {
			return models.Brand{}, errorsResponse(createErrorList, frame.context)
		} else if value, ok := createResult.Value.(models.Brand); ok {
			fmt.Println("Created a new Brand:", createResult.Value)

			return value, nil
		} else {
			err := errorResponse(errors.New("Error trying to create a brand"), frame.context)

			return models.Brand{}, err
		}
	}

	if len(getErrorList) > 0 {
		return models.Brand{}, errorsResponse(getErrorList, frame.context)
	} else if value, ok := getResult.Value.(models.Brand); ok {
		return value, nil
	} else {
		err := errorResponse(errors.New("Error trying to get a brand from the database"), frame.context)

		return models.Brand{}, err
	}
}

func validateRating(from uint8, to uint8, frame *frame) error {
	if (frame.request.Rating <= from) || (frame.request.Rating >= to) {
		errorMessage := fmt.Sprintf("Error validating rating: %v is not in range(%v, %v)",
			frame.request.Rating,
			from,
			to)

		return responses.ErrorResponse(http.StatusUnprocessableEntity, errorMessage, frame.context)
	}

	return nil
}
