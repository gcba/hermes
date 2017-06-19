package responses

type Post struct {
	Meta Meta `json:"meta"`
}

func PostResponse(status int, context echo.Context) error {
	response := Post{
		Meta: metas[status]
	}

  	return context.JSON(status, &response)
}
