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
	request, err := parser.Parse(echoContext)

	if err != nil {
		return err
	}

	if !echoContext.Response().Committed {
		db := database.GetReadDB()
		defer db.Close()

		baseContext := echoContext.Request().Context()
		loadedContext := context.WithValue(baseContext, schema.DB, db)
		response := schema.Schema.Exec(loadedContext, request.Query, "", request.Variables)

		return responses.PostResponse(echoContext, response)
	}

	return nil
}
