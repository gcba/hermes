package responses

import (
	"github.com/labstack/echo"
	"net/http"

	"hermes/responses"
)

type Post struct {
	Meta responses.Status `json:"meta"`
}

func PostResponse(context echo.Context) error {
	if !context.Response().Committed {
		response := Post{Meta: responses.Statuses[http.StatusCreated]}

		return context.JSON(http.StatusCreated, &response)
	}

	return nil
}
