package controller

import (
	"errors"
	"fmt"
	"net/http"

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

	context.Response().Header().Set(echo.HeaderAllow, "OPTIONS POST")

	return responses.OptionsResponse(endpoints, context)
}

// PostRatings saves a new rating to the database
func PostRatings(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	readDB := database.GetReadDB()
	writeDB := database.GetWriteDB()
	dbs := &databases{read: readDB, write: writeDB}
	frame := &frame{request: request, context: context}

	defer readDB.Close()
	defer writeDB.Close()

	if err := newRating(dbs, frame); err != nil {
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

	if len(errorList) > 0 {
		return errorsResponse(errorList, frame.context)
	} else if value, ok := result.Value.(models.Message); ok {
		fmt.Println("Created a new Message:", value)

		return nil
	} else {
		return errorResponse(errors.New("Error trying to create a message"), frame.context)
	}
}

/*
*
* Rating
*
 */
func newRating(dbs *databases, frame *frame) error {
	app, appErr := getApp(dbs.read, frame)

	if appErr != nil {
		return appErr
	}

	platform, platformErr := getPlatform(dbs.read, frame)

	if platformErr != nil {
		return platformErr
	}

	rangeRecord, rangeErr := getRange(dbs.read, frame)

	if rangeErr != nil {
		return rangeErr
	}

	if err := validateRating(rangeRecord.From, rangeRecord.To, frame); err != nil {
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
		BrowserVersion:  frame.request.Browser.Version,
		HasMessage:      hasMessage,
		AppID:           app.ID,
		RangeID:         rangeRecord.ID,
		PlatformID:      platform.ID}

	if hasAppUser(frame.request) {
		if err := attachAppUser(rating, dbs, frame); err != nil {
			return err
		}
	}

	if hasDevice(frame.request) {
		if err := attachDevice(rating, &platform, dbs, frame); err != nil {
			return err
		}
	}

	if hasBrowser(frame.request) {
		if err := attachBrowser(rating, dbs, frame); err != nil {
			return err
		}
	}

	result := models.CreateRating(rating, dbs.write)
	errorList := result.GetErrors()

	if len(errorList) > 0 {
		return errorsResponse(errorList, frame.context)
	} else if value, ok := result.Value.(models.Rating); ok {
		fmt.Println("Created a new Rating:", value)

		if hasMessage {
			if err := newMessage(value.ID, dbs.write, frame); err != nil {
				return err
			}
		}

		return nil
	} else {
		return errorResponse(errors.New("Error trying to create a rating"), frame.context)
	}
}
