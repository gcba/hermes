package responses

import (
	"github.com/labstack/echo"
	"net/http"
)

type Post struct {
	Meta meta `json:"meta"`
}

func PostResponse(context echo.Context) error {
	if !context.Response().Committed {
		response := Post{Meta: metas[http.StatusCreated]}

		return context.JSON(http.StatusCreated, &response)
	}

	return nil
}
