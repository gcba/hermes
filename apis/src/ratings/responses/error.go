package responses

type (
	Error struct {
		Meta   meta     `json:"meta"`
		Errors []string `json:"errors"`
	}
)

func ErrorResponse(status int, errors []string, context echo.Context) error {
	response := Error{
		Meta: metas[status],
		Errors: errors
	}

  	return context.JSON(status, &response)
}