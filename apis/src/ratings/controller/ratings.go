package controller

import (
	"ratings/database"
	"ratings/models"
	"ratings/parser"
	"ratings/responses"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
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

// OptionsRatings returns the response of the `OPTIONS /ratings` endpoint
func OptionsRatings(context echo.Context) error {
	endpoints := []responses.Endpoint{responses.Endpoints["PostRatings"]}

	context.Response().Header().Set(echo.HeaderAllow, "OPTIONS, POST")

	return responses.OptionsResponse(endpoints, context)
}

// PostRatings saves a new rating to the database
func PostRatings(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		context.Logger().Error("Error parsing request: " + err.Error())

		return err
	}

	dbs := &databases{read: database.GetWriteDB(), write: database.GetReadDB()}
	frame := &frame{request: request, context: context}

	defer dbs.read.Close()
	defer dbs.write.Close()

	if err := newRating(dbs, frame); err != nil {
		frame.context.Logger().Error("Rating error: " + err.Error())

		return err
	}

	return responses.PostResponse(frame.context)
}

/*
*
* Message
*
 */
func newMessage(rating uint, db *gorm.DB, frame *frame) error {
	message := &models.Message{
		Message:   frame.request.Comment,
		Direction: "in",
		RatingID:  rating}

	result := models.CreateMessage(message, db)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return errorResponse(frame.context)
	}

	if value, ok := result.Value.(*models.Message); ok {
		frame.context.Logger().Info("Created a new Message:", value)

		return nil
	}

	frame.context.Logger().Error("Error creating a new Message: Could not cast to model instance")

	return errorResponse(frame.context)
}

/*
*
* Rating
*
 */
func newRating(dbs *databases, frame *frame) error {
	appChannel := make(chan appResult)
	platformChannel := make(chan platformResult)
	rangeChannel := make(chan rangeResult)

	defer close(appChannel)
	defer close(platformChannel)
	defer close(rangeChannel)

	go getApp(dbs.read, frame, appChannel)
	go getPlatform(dbs.read, frame, platformChannel)
	go getRange(dbs.read, frame, rangeChannel)

	app, appOk := <-appChannel
	platform, platformOk := <-platformChannel
	rangeRecord, rangeOk := <-rangeChannel

	if !appOk || !platformOk || !rangeOk {
		frame.context.Logger().Error("Channel closed")

		return errorResponse(frame.context)
	}

	if app.err != nil {
		frame.context.Logger().Error("Error getting app: " + app.err.Error())

		return app.err
	}

	if platform.err != nil {
		frame.context.Logger().Error("Error getting platform: " + platform.err.Error())

		return platform.err
	}

	if rangeRecord.err != nil {
		frame.context.Logger().Error("Error getting range: " + rangeRecord.err.Error())

		return rangeRecord.err
	}

	if err := validateRating(rangeRecord.value.From, rangeRecord.value.To, frame); err != nil {
		frame.context.Logger().Error("Error validating rating: " + err.Error())

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

	attachDeviceChannel := make(chan error)
	defer close(attachDeviceChannel)

	go attachDevice(rating, platform.value, dbs, frame, attachDeviceChannel)

	deviceError, deviceOk := <-attachDeviceChannel

	if deviceError != nil {
		frame.context.Logger().Error("Error attaching device: " + deviceError.Error())

		return deviceError
	}

	if !deviceOk {
		frame.context.Logger().Error("Channel closed")

		return errorResponse(frame.context)
	}

	if hasAppUser(frame.request) {
		attachAppUserChannel := make(chan error)
		defer close(attachAppUserChannel)

		go attachAppUser(rating, dbs, frame, attachAppUserChannel)

		appUserError, appUserOk := <-attachAppUserChannel

		if appUserError != nil {
			frame.context.Logger().Error("Error attaching appuser: " + appUserError.Error())

			return appUserError
		}

		if !appUserOk {
			frame.context.Logger().Error("Channel closed")

			return errorResponse(frame.context)
		}
	}

	if hasBrowser(frame.request) {
		attachBrowserChannel := make(chan error)
		defer close(attachBrowserChannel)

		go attachBrowser(rating, dbs, frame, attachBrowserChannel)

		browserError, browserOk := <-attachBrowserChannel

		if browserError != nil {
			frame.context.Logger().Error("Error attaching browser: " + browserError.Error())

			return browserError
		}

		if !browserOk {
			frame.context.Logger().Error("Channel closed")

			return errorResponse(frame.context)
		}
	}

	result := models.CreateRating(rating, dbs.write)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return errorResponse(frame.context)
	}

	if value, ok := result.Value.(*models.Rating); ok {
		frame.context.Logger().Info("Created a new Rating:", value)

		if rating.HasMessage {
			if err := newMessage(value.ID, dbs.write, frame); err != nil {
				frame.context.Logger().Error("Error creating a message: " + err.Error())

				return err
			}
		}

		return nil
	}

	frame.context.Logger().Error("Error creating a new Rating: Could not cast to model instance")

	return errorResponse(frame.context)
}
