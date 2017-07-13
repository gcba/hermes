package controller

import (
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

func errorResponse(context echo.Context) error {
	return responses.ErrorResponse(http.StatusInternalServerError, "", context)
}

/*
*
* App
*
 */
func getApp(db *gorm.DB, frame *frame) (*models.App, error) {
	result := models.GetApp(frame.request.App.Key, db)

	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return &models.App{}, errorResponse(frame.context)
	}

	return result.Value.(*models.App), nil
}

/*
*
* Platform
*
 */
func getPlatform(db *gorm.DB, frame *frame) (*models.Platform, error) {
	result := models.GetPlatform(frame.request.Platform.Key, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return &models.Platform{}, errorResponse(frame.context)
	}

	return result.Value.(*models.Platform), nil
}

/*
*
* Range
*
 */
func getRange(db *gorm.DB, frame *frame) (*models.Range, error) {
	result := models.GetRange(frame.request.Range, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return &models.Range{}, errorResponse(frame.context)
	}

	return result.Value.(*models.Range), nil
}

/*
*
* AppUser
*
 */
func hasAppUser(request *parser.Request) bool {
	if request.User == nil {
		return false
	}

	appuser := request.User
	nameLength := len(appuser.Name)
	emailLength := len(appuser.Email)
	mibaIDLength := len(appuser.MiBAID)

	if nameLength == 0 || (emailLength == 0 && mibaIDLength == 0) {
		return false
	}

	return true
}

func getAppUser(dbs *databases, frame *frame) (*models.AppUser, error) {
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
			return &models.AppUser{}, errorResponse(frame.context)
		}

		return createResult.Value.(*models.AppUser), nil
	} else if len(getErrorList) > 0 {
		return &models.AppUser{}, errorResponse(frame.context)
	}

	return getResult.Value.(*models.AppUser), nil
}

func attachAppUser(rating *models.Rating, dbs *databases, frame *frame) error {
	appUser, err := getAppUser(dbs, frame)

	if err == nil {
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
	if request.Browser == nil {
		return false
	}

	return true
}

func getBrowser(dbs *databases, frame *frame) (*models.Browser, error) {
	getResult := models.GetBrowser(frame.request.Browser.Name, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &models.Browser{Name: frame.request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Value == nil {
			return &models.Browser{}, errorResponse(frame.context)
		}

		return createResult.Value.(*models.Browser), nil
	} else if len(getErrorList) > 0 || getResult.Value == nil {
		return &models.Browser{}, errorResponse(frame.context)
	}

	return getResult.Value.(*models.Browser), nil
}

func attachBrowser(rating *models.Rating, dbs *databases, frame *frame) error {
	browser, err := getBrowser(dbs, frame)

	if err == nil {
		rating.BrowserID = browser.ID
		rating.BrowserVersion = frame.request.Browser.Version

		return err
	}

	return nil
}

/*
*
* Device
*
 */
func getDevice(brand *models.Brand, platform *models.Platform, dbs *databases, frame *frame) (*models.Device, error) {
	deviceName := frame.request.Device.Name
	screenWidth := frame.request.Device.Screen.Width
	screenHeight := frame.request.Device.Screen.Height

	if frame.request.Device.Name == "Desktop" {
		deviceName = fmt.Sprintf("Desktop %dx%d", screenWidth, screenHeight)
	}

	getResult := models.GetDevice(deviceName, dbs.read)
	getErrorList := getResult.GetErrors()

	var device *models.Device

	if getResult.RecordNotFound() {
		device = &models.Device{
			Name:         deviceName,
			ScreenWidth:  frame.request.Device.Screen.Width,
			ScreenHeight: frame.request.Device.Screen.Height,
			PPI:          frame.request.Device.Screen.PPI,
			PlatformID:   platform.ID}

		if brand != nil {
			device.BrandID = brand.ID
		}
	}

	if result, ok := getResult.Value.(*models.Device); (ok && brand != nil) && (result.BrandID != brand.ID) {
		checkDeviceName := fmt.Sprintf("%v (%v)", deviceName, brand.Name)
		checkGetResult := models.GetDevice(checkDeviceName, dbs.read)
		checkGetErrorList := checkGetResult.GetErrors()

		fmt.Println("\nFIRST CONTROL POINT")

		if checkGetResult.RecordNotFound() {
			device = &models.Device{
				Name:         checkDeviceName,
				ScreenWidth:  frame.request.Device.Screen.Width,
				ScreenHeight: frame.request.Device.Screen.Height,
				PPI:          frame.request.Device.Screen.PPI,
				PlatformID:   platform.ID,
				BrandID:      brand.ID}
		} else if len(checkGetErrorList) > 0 {
			return &models.Device{}, errorResponse(frame.context)
		}

		return checkGetResult.Value.(*models.Device), nil
	}

	if device != nil {
		createResult := models.CreateDevice(device, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Value == nil {
			return &models.Device{}, errorResponse(frame.context)
		}

		return createResult.Value.(*models.Device), nil
	} else if len(getErrorList) > 0 || getResult.Value == nil {
		return &models.Device{}, errorResponse(frame.context)
	}

	return getResult.Value.(*models.Device), nil
}

func attachDevice(rating *models.Rating, platform *models.Platform, dbs *databases, frame *frame) error {
	var brand *models.Brand
	var brandErr error

	if frame.request.Device.Brand != nil {
		brand, brandErr = getBrand(dbs, frame)

		if brandErr != nil {
			return brandErr
		}
	}

	device, deviceErr := getDevice(brand, platform, dbs, frame)

	if deviceErr == nil {
		rating.DeviceID = device.ID
	}

	return deviceErr
}

/*
*
* Brand
*
 */
func getBrand(dbs *databases, frame *frame) (*models.Brand, error) {
	getResult := models.GetBrand(*frame.request.Device.Brand, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &models.Brand{Name: *frame.request.Device.Brand}
		createResult := models.CreateBrand(brand, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Value == nil {
			return &models.Brand{}, errorResponse(frame.context)
		}

		return createResult.Value.(*models.Brand), nil
	} else if len(getErrorList) > 0 || getResult.Value == nil {
		return &models.Brand{}, errorResponse(frame.context)
	}

	return getResult.Value.(*models.Brand), nil
}

func validateRating(from int8, to int8, frame *frame) error {
	if (frame.request.Rating < from) || (frame.request.Rating > to) {
		errorMessage := fmt.Sprintf("Error validating rating: %v is not in range(%v, %v)",
			frame.request.Rating,
			from,
			to)

		return responses.ErrorResponse(http.StatusUnprocessableEntity, errorMessage, frame.context)
	}

	return nil
}

/*
*
* Message
*
 */
func hasMessage(request *parser.Request) bool {
	result := false

	if len(request.Comment) > 0 {
		result = true
	}

	return result
}
