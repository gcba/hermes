package responses

import (
	"github.com/labstack/echo"
)

type Post struct {
	Meta meta `json:"meta"`
}

func PostResponse(status int, context echo.Context) error {
	response := Post{
		Meta: metas[status]}

	return context.JSON(status, &response)
}
