package controller

import (
	"hermes/database"
	"hermes/models"
	"hermes/ratings/parser"
	"hermes/ratings/responses"

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
		return err
	}

	dbs := &databases{read: database.GetReadDB(), write: database.GetWriteDB()}
	frame := &frame{request: request, context: context}

	defer dbs.read.Close()
	defer dbs.write.Close()

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
		Notified:  false,
		RatingID:  rating}

	result := models.CreateMessage(message, db)
	errorMessage := "Error creating a new Message: "
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		return loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := result.Value.(*models.Message); ok {
		frame.context.Logger().Info("Created a new Message: ", value)

		return nil
	}

	return loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}

/*
*
* Rating
*
 */
func newRating(dbs *databases, frame *frame) error {
	rating, platform, buildErr := buildRating(dbs, frame)
	errorMessage := "Error creating a new Rating: "

	if buildErr != nil {
		return buildErr
	}

	attachErr := attachFields(rating, platform, dbs, frame)

	if attachErr != nil {
		return attachErr
	}

	result := models.CreateRating(rating, dbs.write)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		return loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := result.Value.(*models.Rating); ok {
		frame.context.Logger().Info("Created a new Rating: ", value)

		if rating.HasMessage {
			if err := newMessage(value.ID, dbs.write, frame); err != nil {
				return err
			}
		}

		return nil
	}

	return loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}
