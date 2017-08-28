package responses

import (
	"encoding/json"
	"net/http"

	"hermes/responses"
	"hermes/stats/schema"

	"github.com/labstack/echo"
	"github.com/neelance/graphql-go"
)

func PostResponse(echoContext echo.Context, response *graphql.Response) error {
	if !echoContext.Response().Committed {
		var dataMap map[string]interface{}

		responseMap := map[string]interface{}{}
		status := http.StatusOK

		if len(response.Errors) > 0 {
			if statsError := getCustomError(response); statsError != nil {
				status = statsError.Code
			} else {
				status = http.StatusBadRequest
			}

			responseMap["errors"] = response.Errors
		} else {
			json.Unmarshal(response.Data, &dataMap)

			responseMap["data"] = dataMap
		}

		responseMap["meta"] = map[string]interface{}{
			"code":    status,
			"message": responses.Statuses[status].Message,
		}

		return echoContext.JSON(status, &responseMap)
	}

	return nil
}

func getCustomError(response *graphql.Response) *schema.StatsError {
	for _, err := range response.Errors {
		if err.ResolverError != nil {
			if statsError, castOk := err.ResolverError.(*schema.StatsError); castOk {
				return statsError
			}
		}
	}

	return nil
}
