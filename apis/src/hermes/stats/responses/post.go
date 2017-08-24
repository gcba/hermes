package responses

import (
	"encoding/json"
	"net/http"

	"hermes/responses"

	"github.com/labstack/echo"
	"github.com/neelance/graphql-go"
)

func PostResponse(echoContext echo.Context, response *graphql.Response) error {
	if !echoContext.Response().Committed {
		var dataMap map[string]interface{}

		status := http.StatusOK
		responseMap := map[string]interface{}{}

		metaMap := map[string]interface{}{
			"code":    status,
			"message": responses.Statuses[status].Message,
		}

		if response.Errors != nil {
			status = http.StatusBadRequest
			responseMap["errors"] = response.Errors

			metaMap["code"] = status
			metaMap["message"] = responses.Statuses[status].Message
		}

		json.Unmarshal(response.Data, &dataMap)

		responseMap["meta"] = metaMap
		responseMap["data"] = dataMap

		return echoContext.JSON(status, &responseMap)
	}

	return nil
}
