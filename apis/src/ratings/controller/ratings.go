package controller

import (
	"ratings/database"
	"ratings/models"
	"ratings/parser"
	"ratings/responses"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"fmt"
)

// OptionsRatings returns the response of the `OPTIONS /ratings` endpoint
func OptionsRatings(context echo.Context) error {
	endpoints := []responses.Endpoint{responses.Endpoints["PostRatings"]}

	context.Response().Header().Set(echo.HeaderAllow, "OPTIONS POST")

	return responses.OptionsResponse(endpoints, context)
}

// PostRatings saves a new rating to the database
func PostRatings(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		fmt.Println("\n\nError parsing request: ", err.Error())

		return err
	}

	dbs := &databases{read: database.GetWriteDB(), write: database.GetReadDB()}
	frame := &frame{request: request, context: context}

	defer dbs.read.Close()
	defer dbs.write.Close()

	if err := newRating(dbs, frame); err != nil {
		fmt.Println("\n\nRating error: ", err.Error())

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

	value := result.Value.(*models.Message)

	frame.context.Logger().Info("Created a new Message:", value)

	return nil
}

/*
*
* Rating
*
 */
func newRating(dbs *databases, frame *frame) error {
	app, appErr := getApp(dbs.read, frame)

	if appErr != nil {
		fmt.Println("\n\nError getting app: ", appErr.Error())

		return appErr
	}

	platform, platformErr := getPlatform(dbs.read, frame)

	if platformErr != nil {
		fmt.Println("\n\nError getting platform: ", platformErr.Error())

		return platformErr
	}

	rangeRecord, rangeErr := getRange(dbs.read, frame)

	if rangeErr != nil {
		fmt.Println("\n\nError getting range: ", rangeErr.Error())

		return rangeErr
	}

	if err := validateRating(rangeRecord.From, rangeRecord.To, frame); err != nil {
		fmt.Println("\n\nError validating rating: ", err.Error())

		return err
	}

	hasMessage := false

	if len(frame.request.Comment) > 0 {
		hasMessage = true
	}

	rating := &models.Rating{
		Rating:          frame.request.Rating,
		Description:     frame.request.Description,
		AppVersion:      frame.request.App.Version,
		PlatformVersion: frame.request.Platform.Version,
		HasMessage:      hasMessage,
		AppID:           app.ID,
		RangeID:         rangeRecord.ID,
		PlatformID:      platform.ID}

	if hasAppUser(frame.request) {
		if err := attachAppUser(rating, dbs, frame); err != nil {
			fmt.Println("\n\nError attaching appuser: ", err.Error())

			return err
		}
	}

	if hasDevice(frame.request) {
		if err := attachDevice(rating, platform, dbs, frame); err != nil {
			fmt.Println("\n\nError attaching device: ", err.Error())

			return err
		}
	}

	if hasBrowser(frame.request) {
		if err := attachBrowser(rating, dbs, frame); err != nil {
			fmt.Println("\n\nError attaching browser: ", err.Error())

			return err
		}
	}

	result := models.CreateRating(rating, dbs.write)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Value == nil {
		return errorResponse(frame.context)
	}

	value := result.Value.(*models.Rating)

	if hasMessage {
		if err := newMessage(value.ID, dbs.write, frame); err != nil {
			fmt.Println("\n\nError creating a message: ", err.Error())

			return err
		}
	}

	return nil
}
