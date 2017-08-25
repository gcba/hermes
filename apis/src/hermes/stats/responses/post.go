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

		if len(response.Errors) > 0 {
			errs := getErrors(response)

			status = http.StatusBadRequest
			responseMap["errors"] = errs
			metaMap["code"] = status
			metaMap["message"] = responses.Statuses[status].Message
		} else {
			json.Unmarshal(response.Data, &dataMap)

			responseMap["data"] = dataMap
		}

		responseMap["meta"] = metaMap

		return echoContext.JSON(status, &responseMap)
	}

	return nil
}

func getErrors(response *graphql.Response) []error {
	errs := []error{}

	for _, err := range response.Errors {
		if err.ResolverError != nil {
			err.Message = err.ResolverError.Error()
		}

		errs = append(errs, err)
	}

	return errs
}
