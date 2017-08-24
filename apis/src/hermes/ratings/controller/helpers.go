package controller

import (
	"fmt"
	"net/http"

	"hermes/models"
	"hermes/ratings/parser"

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

	appResult struct {
		value *models.App
		err   error
	}

	platformResult struct {
		value *models.Platform
		err   error
	}

	rangeResult struct {
		value *models.Range
		err   error
	}
)

var (
	invalidValueError  = "Database driver returned invalid value"
	cannotCastError    = "Could not cast to model instance"
	channelClosedError = "Channel closed"
)

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func loggedErrorResponse(tag string, message string, context echo.Context) error {
	context.Logger().Error(tag, message)

	return errorResponse()
}

/*
*
* App
*
 */
func getApp(db *gorm.DB, frame *frame, channel chan appResult) {
	result := models.GetApp(frame.request.App.Key, db)
	errorMessage := "Error getting an App:"
	errorList := result.GetErrors()
	resultStruct := appResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, ok := result.Value.(*models.App); ok {
		resultStruct.value = value

		channel <- resultStruct
		close(channel)

		return
	}

	resultStruct.err = errorResponse()

	channel <- resultStruct
	close(channel)

	frame.context.Logger().Error(errorMessage, cannotCastError)
}

/*
*
* Platform
*
 */
func getPlatform(db *gorm.DB, frame *frame, channel chan platformResult) {
	result := models.GetPlatform(frame.request.Platform.Key, db)
	errorMessage := "Error getting a Platform:"
	errorList := result.GetErrors()
	resultStruct := platformResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, ok := result.Value.(*models.Platform); ok {
		resultStruct.value = value

		channel <- resultStruct
		close(channel)

		return
	}

	resultStruct.err = errorResponse()

	channel <- resultStruct
	close(channel)

	frame.context.Logger().Error(errorMessage, cannotCastError)
}

/*
*
* Range
*
 */
func getRange(db *gorm.DB, frame *frame, channel chan rangeResult) {
	result := models.GetRange(frame.request.Range, db)
	errorMessage := "Error getting a Range:"
	errorList := result.GetErrors()
	resultStruct := rangeResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, ok := result.Value.(*models.Range); ok {
		resultStruct.value = value

		channel <- resultStruct
		close(channel)

		return
	}

	resultStruct.err = errorResponse()

	channel <- resultStruct
	close(channel)

	frame.context.Logger().Error(errorMessage, cannotCastError)
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
	var errorMessage = "Could not get an AppUser: "
	var hasMibaID = len(frame.request.User.MiBAID) > 0

	if hasMibaID {
		getResult = models.GetAppUser(frame.request.User.MiBAID, dbs.read)
	} else {
		getResult = models.GetAppUserByEmail(frame.request.User.Email, dbs.read)
	}

	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		var appuser *models.AppUser

		if hasMibaID {
			appuser = &models.AppUser{
				Name:   frame.request.User.Name,
				MiBAID: &frame.request.User.MiBAID}
		} else {
			appuser = &models.AppUser{
				Name:  frame.request.User.Name,
				Email: &frame.request.User.Email}
		}

		createResult := models.CreateAppUser(appuser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.AppUser{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
		}

		if value, ok := createResult.Value.(*models.AppUser); ok {
			frame.context.Logger().Info("Created a new AppUser: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.AppUser{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := getResult.Value.(*models.AppUser); ok {
		return value, nil
	}

	return &models.AppUser{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}

func attachAppUser(rating *models.Rating, dbs *databases, frame *frame, channel chan error) {
	appUser, err := getAppUser(dbs, frame)

	if err == nil {
		rating.AppUserID = appUser.ID
	}

	channel <- err
	close(channel)
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
	errorMessage := "Could not get a Browser: "
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &models.Browser{Name: frame.request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Browser{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
		}

		if value, ok := createResult.Value.(*models.Browser); ok {
			frame.context.Logger().Info("Created a new Browser: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Browser{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := getResult.Value.(*models.Browser); ok {
		return value, nil
	}

	return &models.Browser{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}

func attachBrowser(rating *models.Rating, dbs *databases, frame *frame, channel chan error) {
	browser, err := getBrowser(dbs, frame)

	if err == nil {
		rating.BrowserID = browser.ID
		rating.BrowserVersion = frame.request.Browser.Version
	}

	channel <- err
	close(channel)
}

/*
*
* Device
*
 */
func getDevice(brand *models.Brand, platform *models.Platform, dbs *databases, frame *frame) (*models.Device, error) {
	var device *models.Device
	var errorMessage = "Could not get a Device: "

	deviceName := frame.request.Device.Name
	screenWidth := frame.request.Device.Screen.Width
	screenHeight := frame.request.Device.Screen.Height

	if frame.request.Device.Name == "Desktop" {
		deviceName = fmt.Sprintf("Desktop %dx%d", screenWidth, screenHeight)
	}

	getResult := models.GetDevice(deviceName, dbs.read)
	getErrorList := getResult.GetErrors()

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
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Device{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if result, ok := getResult.Value.(*models.Device); (ok && brand != nil) && (result.BrandID != brand.ID) {
		checkDeviceName := fmt.Sprintf("%v (%v)", deviceName, brand.Name)
		checkGetResult := models.GetDevice(checkDeviceName, dbs.read)
		checkGetErrorList := checkGetResult.GetErrors()

		if checkGetResult.RecordNotFound() {
			device = &models.Device{
				Name:         checkDeviceName,
				ScreenWidth:  frame.request.Device.Screen.Width,
				ScreenHeight: frame.request.Device.Screen.Height,
				PPI:          frame.request.Device.Screen.PPI,
				PlatformID:   platform.ID,
				BrandID:      brand.ID}
		} else if len(checkGetErrorList) > 0 || checkGetResult.Error != nil || checkGetResult.Value == nil {
			return &models.Device{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
		} else {
			if value, ok := checkGetResult.Value.(*models.Device); ok {
				return value, nil
			}

			return &models.Device{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
		}
	}

	if device != nil {
		createResult := models.CreateDevice(device, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Device{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
		}

		if value, ok := createResult.Value.(*models.Device); ok {
			frame.context.Logger().Info("Created a new Device: ", value)

			return value, nil
		}

		return &models.Device{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
	}

	if value, ok := getResult.Value.(*models.Device); ok {
		return value, nil
	}

	return &models.Device{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}

func attachDevice(rating *models.Rating, platform *models.Platform, dbs *databases, frame *frame, channel chan error) {
	var brand *models.Brand
	var brandErr error

	if frame.request.Device.Brand != nil {
		brand, brandErr = getBrand(dbs, frame)

		if brandErr != nil {
			channel <- brandErr
			close(channel)

			return
		}
	}

	device, deviceErr := getDevice(brand, platform, dbs, frame)

	if deviceErr == nil {
		rating.DeviceID = device.ID
	}

	channel <- deviceErr
	close(channel)
}

/*
*
* Brand
*
 */
func getBrand(dbs *databases, frame *frame) (*models.Brand, error) {
	getResult := models.GetBrand(*frame.request.Device.Brand, dbs.read)
	errorMessage := "Could not get a Brand: "
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &models.Brand{Name: *frame.request.Device.Brand}
		createResult := models.CreateBrand(brand, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Brand{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
		}

		if value, ok := createResult.Value.(*models.Brand); ok {
			frame.context.Logger().Info("Created a new Brand: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Brand{}, loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := getResult.Value.(*models.Brand); ok {
		return value, nil
	}

	return &models.Brand{}, loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}

/*
*
* Rating
*
 */
func buildRating(dbs *databases, frame *frame) (*models.Rating, *models.Platform, error) {
	appChannel := make(chan appResult)
	platformChannel := make(chan platformResult)
	rangeChannel := make(chan rangeResult)

	go getApp(dbs.read, frame, appChannel)
	go getPlatform(dbs.read, frame, platformChannel)
	go getRange(dbs.read, frame, rangeChannel)

	app, appOk := <-appChannel
	platform, platformOk := <-platformChannel
	rangeRecord, rangeOk := <-rangeChannel

	if !appOk || !platformOk || !rangeOk {
		return nil, nil, loggedErrorResponse("", channelClosedError, frame.context)
	}

	if app.err != nil {
		return nil, nil, app.err
	}

	if platform.err != nil {
		return nil, nil, platform.err
	}

	if rangeRecord.err != nil {
		return nil, nil, rangeRecord.err
	}

	if err := validateRating(rangeRecord.value.From, rangeRecord.value.To, frame); err != nil {
		return nil, nil, err
	}

	rating := &models.Rating{
		Rating:          frame.request.Rating,
		Description:     frame.request.Description,
		AppVersion:      frame.request.App.Version,
		PlatformVersion: frame.request.Platform.Version,
		HasMessage:      hasMessage(frame.request),
		AppID:           app.value.ID,
		RangeID:         rangeRecord.value.ID,
		PlatformID:      platform.value.ID}

	return rating, platform.value, nil
}

func validateRating(from int8, to int8, frame *frame) error {
	if (frame.request.Rating < from) || (frame.request.Rating > to) {
		errorMessage := fmt.Sprintf("Error validating rating: %v is not in range (%v, %v)",
			frame.request.Rating,
			from,
			to)

		frame.context.Logger().Error(errorMessage)

		return echo.NewHTTPError(http.StatusUnprocessableEntity, []string{errorMessage})
	}

	return nil
}

func attachFields(rating *models.Rating, platform *models.Platform, dbs *databases, frame *frame) error {
	attachDeviceChannel := make(chan error)

	go attachDevice(rating, platform, dbs, frame, attachDeviceChannel)

	deviceError, deviceOk := <-attachDeviceChannel

	if deviceError != nil {
		return loggedErrorResponse("", "Error attaching Device", frame.context)
	}

	if !deviceOk {
		return loggedErrorResponse("", channelClosedError, frame.context)
	}

	if hasAppUser(frame.request) {
		attachAppUserChannel := make(chan error)

		go attachAppUser(rating, dbs, frame, attachAppUserChannel)

		appUserError, appUserOk := <-attachAppUserChannel

		if appUserError != nil {
			return loggedErrorResponse("", "Error attaching AppUser", frame.context)
		}

		if !appUserOk {
			return loggedErrorResponse("", channelClosedError, frame.context)
		}
	}

	if hasBrowser(frame.request) {
		attachBrowserChannel := make(chan error)

		go attachBrowser(rating, dbs, frame, attachBrowserChannel)

		browserError, browserOk := <-attachBrowserChannel

		if browserError != nil {
			return loggedErrorResponse("", "Error attaching Browser", frame.context)
		}

		if !browserOk {
			return loggedErrorResponse("", channelClosedError, frame.context)
		}
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
