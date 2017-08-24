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

	if !echoContext.Response().Committed {
		db := database.GetReadDB()
		defer db.Close()

		currentContext := echoContext.Request().Context()
		loadedContext := context.WithValue(currentContext, schema.DB, db)
		variables := map[string]interface{}{
			"field": mapStruct(&request.Variables.Field),
		}

		response := schema.Schema.Exec(loadedContext, request.Query, "", variables)
		// response := schema.Schema.Exec(loadedContext, request.Query, "", request.Variables)

		return echoContext.JSON(http.StatusOK, &response)
	}

	return nil
}

func mapStruct(field *parser.Field) map[string]interface{} {
	var fieldMap map[string]interface{}

	if field.Next != nil {
		nextMap := map[string]interface{}{
			"next": map[string]interface{}{
				"condition": field.Next.Condition,
				"field":     mapStruct(field.Next.Field),
			},
		}

		fieldMap["next"] = nextMap
	} else {
		fieldMap = map[string]interface{}{
			"name":     field.Name,
			"operator": field.Operator,
			"value":    field.Value,
		}
	}

	return fieldMap
}
