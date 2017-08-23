package controller

import (
	"net/http"
	"strings"

	"hermes/stats/parser"
	"hermes/stats/schema"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/linkosmos/mapop"
)

// PostStats is the main GraphQL controller
func PostStats(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	if !context.Response().Committed {
		variables := mapop.MapKeys(strings.ToLower, structs.Map(&request.Variables))
		response := schema.Schema.Exec(context.Request().Context(), request.Query, "", variables)

		return context.JSON(http.StatusOK, &response)
	}

	return nil
}
