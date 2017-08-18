package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

type Post struct {
	Meta meta `json:"meta"`
}

func PostResponse(context echo.Context) error {
	if !context.Response().Committed {
		response := Post{Meta: metas[http.StatusOK]}

		return context.JSON(http.StatusOK, &response)
	}

	return nil
}
