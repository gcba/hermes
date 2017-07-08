package responses

import (
	"github.com/labstack/echo"
	"net/http"
)

type Post struct {
	Meta meta `json:"meta"`
}

func PostResponse(context echo.Context) error {
	response := Post{Meta: metas[http.StatusCreated]}

	if !context.Response().Committed {
		return context.JSON(http.StatusCreated, &response)
	}

	return nil
}
