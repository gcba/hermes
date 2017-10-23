package controller

import (
	"context"

	"hermes/database"
	"hermes/stats/parser"
	"hermes/stats/responses"
	"hermes/stats/schema"

	"github.com/labstack/echo"
)

// PostStats is the GraphQL controller
func PostStats(echoContext echo.Context) error {
	request, parseErr := parser.Parse(echoContext)

	if parseErr != nil {
		return parseErr
	}

	if !echoContext.Response().Committed {
		db, dbErr := database.GetReadDB()

		if dbErr != nil {
			return dbErr
		}

		defer db.Close()

		baseContext := echoContext.Request().Context()
		loadedContext := context.WithValue(baseContext, schema.DB, db)
		response := schema.Schema.Exec(loadedContext, request.Query, "", request.Variables)

		return responses.PostResponse(echoContext, response)
	}

	return nil
}
