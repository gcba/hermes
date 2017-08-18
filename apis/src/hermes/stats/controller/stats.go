package controller

import (
	"context"
	"net/http"

	"hermes/database"
	"hermes/stats/parser"
	"hermes/stats/schema"

	"github.com/labstack/echo"
)

// PostStats is the main GraphQL controller
func PostStats(echoContext echo.Context) error {
	request, err := parser.Parse(echoContext)

	if err != nil {
		return err
	}

	db := database.GetReadDB()

	defer db.Close()

	if !echoContext.Response().Committed {
		currentContext := echoContext.Request().Context()
		loadedContext := context.WithValue(currentContext, schema.DB, db)
		response := schema.Schema.Exec(loadedContext, request.Query, "", request.Variables)

		return echoContext.JSON(http.StatusOK, &response)
	}

	return nil
}
