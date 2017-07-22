package controller

import (
	"ratings/database"
	"ratings/models"
	"ratings/parser"
	"ratings/responses"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
	rating, platform, err := buildRating(dbs, frame)

	if err != nil {
		return errorResponse(frame.context)
	}

	attachDeviceChannel := make(chan error)
	defer close(attachDeviceChannel)

	go attachDevice(rating, platform, dbs, frame, attachDeviceChannel)

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
