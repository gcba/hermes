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

	records struct {
		rating   *models.Rating
		app      *models.App
		platform *models.Platform
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

	deviceResult struct {
		value *models.Device
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
	errorMessage := "Error getting an App:"
	result := models.GetApp(frame.request.App.Key, db)
	errorList := result.GetErrors()
	resultStruct := appResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, castOk := result.Value.(*models.App); castOk {
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
	errorMessage := "Error getting a Platform:"
	result := models.GetPlatform(frame.request.Platform.Key, db)
	errorList := result.GetErrors()
	resultStruct := platformResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, castOk := result.Value.(*models.Platform); castOk {
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
	errorMessage := "Error getting a Range:"
	result := models.GetRange(frame.request.Range, db)
	errorList := result.GetErrors()
	resultStruct := rangeResult{}

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		resultStruct.err = errorResponse()

		channel <- resultStruct
		close(channel)

		frame.context.Logger().Error(errorMessage, invalidValueError)

		return
	}

	if value, castOk := result.Value.(*models.Range); castOk {
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

func getAppUser(records *records, dbs *databases, frame *frame, deviceChan chan deviceResult) (*models.AppUser, error) {
	var getResult *gorm.DB

	getErrorMessage := "Could not get an AppUser: "
	createErrorMessage := "Could not create an AppUser: "
	hasMibaID := len(frame.request.User.MiBAID) > 0
	hasEmail := len(frame.request.User.Email) > 0

	if hasMibaID {
		getResult = models.GetAppUser(frame.request.User.MiBAID, dbs.read)
	} else {
		getResult = models.GetAppUserByEmail(frame.request.User.Email, dbs.read)
	}

	getErrorList := getResult.GetErrors()
	device, deviceOk := <-deviceChan

	if getResult.RecordNotFound() {
		if !deviceOk {
			return &models.AppUser{}, loggedErrorResponse(createErrorMessage, "Could not get a Device", frame.context)
		}

		if device.err != nil {
			return &models.AppUser{}, loggedErrorResponse(createErrorMessage, device.err.Error(), frame.context)
		}

		appuser := &models.AppUser{
			Name:      frame.request.User.Name,
			Apps:      []models.App{*records.app},
			Platforms: []models.Platform{*records.platform},
			Devices:   []models.Device{*device.value}}

		if hasMibaID {
			appuser.MiBAID = &frame.request.User.MiBAID
		}

		if hasEmail {
			appuser.Email = &frame.request.User.Email
		}

		createResult := models.CreateAppUser(appuser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.AppUser{}, loggedErrorResponse(createErrorMessage, invalidValueError, frame.context)
		}

		if value, castOk := createResult.Value.(*models.AppUser); castOk {
			frame.context.Logger().Info("Created a new AppUser: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.AppUser{}, loggedErrorResponse(getErrorMessage, invalidValueError, frame.context)
	}

	if value, castOk := getResult.Value.(*models.AppUser); castOk {
		return value, nil
	}

	return &models.AppUser{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
}

func attachAppUser(records *records, dbs *databases, frame *frame, deviceChan chan deviceResult, errorChan chan error) {
	appUser, err := getAppUser(records, dbs, frame, deviceChan)

	if err == nil {
		records.rating.AppUserID = appUser.ID
	}

	errorChan <- err
	close(errorChan)
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
	getErrorMessage := "Could not get a Browser: "
	createErrorMessage := "Could not create a Browser: "
	getResult := models.GetBrowser(frame.request.Browser.Name, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		browser := &models.Browser{Name: frame.request.Browser.Name}
		createResult := models.CreateBrowser(browser, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Browser{}, loggedErrorResponse(createErrorMessage, invalidValueError, frame.context)
		}

		if value, castOk := createResult.Value.(*models.Browser); castOk {
			frame.context.Logger().Info("Created a new Browser: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Browser{}, loggedErrorResponse(getErrorMessage, invalidValueError, frame.context)
	}

	if value, castOk := getResult.Value.(*models.Browser); castOk {
		return value, nil
	}

	return &models.Browser{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
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

	getErrorMessage := "Could not get a Device: "
	createErrorMessage := "Could not create a Device: "
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
			device.BrandID = &brand.ID
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Device{}, loggedErrorResponse(getErrorMessage, invalidValueError, frame.context)
	}

	if result, castOk := getResult.Value.(*models.Device); (castOk && brand != nil) && (result.BrandID != &brand.ID) {
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
				BrandID:      &brand.ID}
		} else if len(checkGetErrorList) > 0 || checkGetResult.Error != nil || checkGetResult.Value == nil {
			return &models.Device{}, loggedErrorResponse(getErrorMessage, invalidValueError, frame.context)
		} else {
			if value, checkCastOk := checkGetResult.Value.(*models.Device); checkCastOk {
				return value, nil
			}

			return &models.Device{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
		}
	} else if !castOk {
		return &models.Device{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
	}

	if device != nil {
		createResult := models.CreateDevice(device, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Device{}, loggedErrorResponse(createErrorMessage, invalidValueError, frame.context)
		}

		if value, castOk := createResult.Value.(*models.Device); castOk {
			frame.context.Logger().Info("Created a new Device: ", value)

			return value, nil
		}

		return &models.Device{}, loggedErrorResponse(createErrorMessage, cannotCastError, frame.context)
	}

	if value, castOk := getResult.Value.(*models.Device); castOk {
		return value, nil
	}

	return &models.Device{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
}

func attachDevice(records *records, dbs *databases, frame *frame, modelChan chan deviceResult, errorChan chan error) {
	var brand *models.Brand
	var brandErr error

	resultStruct := deviceResult{}

	if frame.request.Device.Brand != nil {
		brand, brandErr = getBrand(dbs, frame)

		if brandErr != nil {
			resultStruct.err = brandErr

			modelChan <- resultStruct
			close(modelChan)
			errorChan <- brandErr
			close(errorChan)

			return
		}
	}

	device, deviceErr := getDevice(brand, records.platform, dbs, frame)

	if deviceErr == nil {
		resultStruct.value = device
		records.rating.DeviceID = device.ID
	} else {
		resultStruct.err = deviceErr
	}

	modelChan <- resultStruct
	close(modelChan)
	errorChan <- deviceErr
	close(errorChan)
}

/*
 *
 * Brand
 *
 */
func getBrand(dbs *databases, frame *frame) (*models.Brand, error) {
	getErrorMessage := "Could not get a Brand: "
	createErrorMessage := "Could not create a Brand: "
	getResult := models.GetBrand(*frame.request.Device.Brand, dbs.read)
	getErrorList := getResult.GetErrors()

	if getResult.RecordNotFound() {
		brand := &models.Brand{Name: *frame.request.Device.Brand}
		createResult := models.CreateBrand(brand, dbs.write)
		createErrorList := createResult.GetErrors()

		if len(createErrorList) > 0 || createResult.Error != nil || createResult.Value == nil {
			dbs.write.Rollback()

			return &models.Brand{}, loggedErrorResponse(createErrorMessage, invalidValueError, frame.context)
		}

		if value, castOk := createResult.Value.(*models.Brand); castOk {
			frame.context.Logger().Info("Created a new Brand: ", value)

			return value, nil
		}
	} else if len(getErrorList) > 0 || getResult.Error != nil || getResult.Value == nil {
		return &models.Brand{}, loggedErrorResponse(getErrorMessage, invalidValueError, frame.context)
	}

	if value, castOk := getResult.Value.(*models.Brand); castOk {
		return value, nil
	}

	return &models.Brand{}, loggedErrorResponse(getErrorMessage, cannotCastError, frame.context)
}

/*
 *
 * Rating
 *
 */
func buildRating(records *records, dbs *databases, frame *frame) error {
	appChannel := make(chan appResult, 1)
	platformChannel := make(chan platformResult, 1)
	rangeChannel := make(chan rangeResult, 1)

	go getApp(dbs.read, frame, appChannel)
	go getPlatform(dbs.read, frame, platformChannel)
	go getRange(dbs.read, frame, rangeChannel)

	app, appOk := <-appChannel
	platform, platformOk := <-platformChannel
	rangeRecord, rangeOk := <-rangeChannel

	if !appOk || !platformOk || !rangeOk {
		return loggedErrorResponse("", channelClosedError, frame.context)
	}

	if app.err != nil {
		return app.err
	}

	if platform.err != nil {
		return platform.err
	}

	if rangeRecord.err != nil {
		return rangeRecord.err
	}

	if err := validateRating(rangeRecord.value.From, rangeRecord.value.To, frame); err != nil {
		return err
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

	records.rating = rating
	records.app = app.value
	records.platform = platform.value

	return nil
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

func attachFields(records *records, dbs *databases, frame *frame) error {
	deviceChannel := make(chan deviceResult, 1)
	attachDeviceChannel := make(chan error, 1)

	go attachDevice(records, dbs, frame, deviceChannel, attachDeviceChannel)

	deviceError, deviceOk := <-attachDeviceChannel

	if deviceError != nil {
		return loggedErrorResponse("", "Error attaching Device", frame.context)
	}

	if !deviceOk {
		return loggedErrorResponse("", channelClosedError, frame.context)
	}

	if hasAppUser(frame.request) {
		attachAppUserChannel := make(chan error, 1)

		go attachAppUser(records, dbs, frame, deviceChannel, attachAppUserChannel)

		appUserError, appUserOk := <-attachAppUserChannel

		if appUserError != nil {
			return loggedErrorResponse("", "Error attaching AppUser", frame.context)
		}

		if !appUserOk {
			return loggedErrorResponse("", channelClosedError, frame.context)
		}
	}

	if hasBrowser(frame.request) {
		attachBrowserChannel := make(chan error, 1)

		go attachBrowser(records.rating, dbs, frame, attachBrowserChannel)

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
