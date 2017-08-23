package controller

import (
	"net/http"

	"hermes/stats/parser"
	"hermes/stats/schema"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
)

// PostStats is the main GraphQL controller
func PostStats(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	if !context.Response().Committed {
		response := schema.Schema.Exec(context.Request().Context(), request.Query, "", structs.Map(request.Variables))

		return context.JSON(http.StatusOK, &response)
	}

	return nil
}
