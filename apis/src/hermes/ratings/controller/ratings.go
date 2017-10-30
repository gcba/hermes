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

	return responses.OptionsResponse(endpoints, "OPTIONS, POST", context)
}

// PostRatings saves a new rating to the database
func PostRatings(context echo.Context) error {
	request, parseErr := parser.Parse(context)

	if parseErr != nil {
		return parseErr
	}

	if !context.Response().Committed {
		readDB, readDBErr := database.GetReadDB()

		if readDBErr != nil {
			return readDBErr
		}

		defer readDB.Close()

		writeDB, writeDBErr := database.GetWriteDB()

		if writeDBErr != nil {
			return writeDBErr
		}

		defer writeDB.Close()

		tx := writeDB.Begin()
		dbs := &databases{read: readDB, write: tx}
		frame := &frame{request: request, context: context}

		if ratingErr := newRating(dbs, frame); ratingErr != nil {
			return ratingErr
		}

		return responses.PostResponse(frame.context)
	}

	return nil
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
		Status:    0,
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
	records := &records{}
	buildErr := buildRating(records, dbs, frame)
	errorMessage := "Error creating a new Rating: "

	if buildErr != nil {
		return buildErr
	}

	attachErr := attachFields(records, dbs, frame)

	if attachErr != nil {
		return attachErr
	}

	result := models.CreateRating(records.rating, dbs.write)
	errorList := result.GetErrors()

	if len(errorList) > 0 || result.Error != nil || result.Value == nil {
		dbs.write.Rollback()

		return loggedErrorResponse(errorMessage, invalidValueError, frame.context)
	}

	if value, ok := result.Value.(*models.Rating); ok {
		frame.context.Logger().Info("Created a new Rating: ", value)

		if records.rating.HasMessage {
			if err := newMessage(value.ID, dbs.write, frame); err != nil {
				dbs.write.Rollback()

				return err
			}
		}

		dbs.write.Commit()

		return nil
	}

	return loggedErrorResponse(errorMessage, cannotCastError, frame.context)
}
