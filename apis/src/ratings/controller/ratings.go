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
func PostRating(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	readDb := GetReadDB()
	defer readDb.Close()
	writeDb := GetWriteDB()
	defer writeDb.Close()

	appRecord, rangeRecord, platformRecord := getBaseFields(request, readDb)

	if rating, ok := newRating(request, appRecord, rangeRecord, platformRecord, readDb, writeDb)

	if !ok {
		// TODO: Dispatch error response

		return errors.New("Rating creation failed")
	}

	return context.JSON(http.StatusOK, &rating)
}

func OptionsRating(context echo.Context) error {
	meta := responses.Meta{
		Code: 200,
		Message: "Request completed successfully"
	}

	headers := responses.Header{
		ContentType: "application/json; charset=utf-8"
	}

	postRatings := responses.Endpoint{
		Method: "POST",
		Path: "/ratings",
		Headers: headers
	}

	response := responses.Options{
		Meta:  meta
		Endpoints: []Endpoint{postRatings}
	}

  	return context.JSON(http.StatusOK, &response)
}

func getBaseFields(request *parser.Request, db *gorm.DB) (models.App, models.Range, models.Platform) {
	appRecord, appErr := models.GetApp(request.App.Key, db)

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

func getBrowser(request *parser.Request, readDB *gorm.DB, writeDB *gorm.DB) (models.Browser, bool) { // TODO: Move this to models
	getResult := models.GetBrowser(request.Browser.Name, readDB)

	if getResult.RecordNotFound() {
		browser := Browser{Name: request.Browser.Name}
		createResult := models.CreateBrowser(&browser, writeDB)

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

func getBrand(request *parser.Request, readDB *gorm.DB, writeDB *gorm.DB) (models.Brand, bool) { // TODO: Move this to models
	getResult := models.GetBrand(request.Browser.Name, readDB)

	if getResult.RecordNotFound() {
		brand := Brand{Name: request.Brand}
		createResult := models.CreateBrand(&brand, writeDB)

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

func getDevice(request *parser.Request, brandRecord *models.Brand, platformRecord *models.Platform, readDB *gorm.DB, writeDB *gorm.DB) (models.Device, bool) { // TODO: Move this to models
	getResult := models.GetDevice(request.Device.Name, request.Brand, readDB)

	if getResult.RecordNotFound() {
		device := Device{
				Name: request.Device.Name,
				ScreenWidth: request.Device.Screen.Width,
				ScreenHeight: request.Device.Screen.Height,
				PPI: request.Device.Screen.PPI,
				BrandID: brandRecord.ID,
				PlatformID: platformRecord.ID
			}

		createResult := models.CreateDevice(&device, writeDB)

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

func getAppUser(request *parser.Request, readDB *gorm.DB, writeDB *gorm.DB) (models.AppUser, bool) { // TODO: Move this to models
	getResult := models.GetAppUser(request.User.MiBAID, readDB)

	if getResult.RecordNotFound() {
		appuser := AppUser{
			Name: request.User.Name,
			Email: request.User.Email,
			MiBAID: request.User.MiBAID
		}

		createResult := models.CreateAppUser(&appuser, writeDB)

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

func newRating(request *parser.Request, appRecord *models.App, rangeRecord *models.Range, platformRecord *models.Platform, writeDb *gorm.DB, readDb *gorm.DB) (models.Rating, bool) {
	hasMessage := false

	if len(request.Comment) > 0 {
		hasMessage = true
	}

	rating := Rating{
		Rating: request.Rating,
		Description: request.Description,
		AppVersion: request.App.Version,
		PlatformVersion: request.Platform.Version,
		BrowserVersion: request.Browser.Version,
		HasMessage: hasMessage,
		AppID: appRecord.ID,
		RangeID: rangeRecord.ID,
		PlatformID: platformRecord.ID
	}

	if hasAppUser(request) {
		appUser, ok := getAppUser(request, readDb)

		if !ok {
			// TODO: Handle error
			fmt.Println("Could not get AppUser")

			return Rating{}, false
		}

		rating.AppUserID = appUser.ID
	}

	if hasBrowser(request) {
		browser, ok := getBrowser(request, readDb)

		if !ok {
			// TODO: Handle error
			fmt.Println("Could not get Browser")

			return Rating{}, false
		}

		rating.BrowserID = browser.ID
	}

	if hasDevice(request) {
		brand, brandOk := getBrand(request, readDb)

		if !brandOk {
			// TODO: Handle error
			fmt.Println("Could not get Brand")

			return Rating{}, false
		}

		device, deviceOk := getDevice(request, brand, platformRecord, readDb)

		if !deviceOk {
			// TODO: Handle error
			fmt.Println("Could not get Device")

			return Rating{}, false
		}

		rating.DeviceID = device.ID
	}

	createResult := models.CreateRating(&rating, writeDb)

	if len(createResult.GetErrors()) > 0 {
		// TODO: Handle errors

		for value, err := range getResult.GetErrors() {
			fmt.Println("Error creating a Rating:", value)
		}
	}

	fmt.Println("Created a new Rating:", createResult.Value)

	return createResult.Value, true
}