package responses

type (
	Error struct {
		Meta   meta     `json:"meta"`
		Errors []string `json:"errors"`
	}
)

func ErrorResponse(status int, singleError string, context echo.Context) error {
	return ErrorsResponse(status, []string{singleError}, context)
}

func ErrorsResponse(status int, errors []string, context echo.Context) error {
	response := Error{
		Meta: metas[status],
		Errors: errors
	}

  	return context.JSON(status, &response)
}