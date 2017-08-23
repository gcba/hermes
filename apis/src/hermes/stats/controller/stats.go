package controller

import (
	"context"
	"net/http"
	"strings"

	"hermes/database"
	"hermes/stats/parser"
	"hermes/stats/schema"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/linkosmos/mapop"
)

// PostStats is the main GraphQL controller
func PostStats(echoContext echo.Context) error {
	request, err := parser.Parse(echoContext)

	if err != nil {
		return err
	}

	if !echoContext.Response().Committed {
		db := database.GetReadDB()
		defer db.Close()

		currentContext := echoContext.Request().Context()
		loadedContext := context.WithValue(currentContext, schema.DB, db)
		variables := mapop.MapKeys(strings.ToLower, structs.Map(&request.Variables))
		response := schema.Schema.Exec(loadedContext, request.Query, "", variables)

		return echoContext.JSON(http.StatusOK, &response)
	}

	return nil
}
