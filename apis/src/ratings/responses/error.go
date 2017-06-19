package responses

type (
	ErrorType struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	}

	Error struct {
		Meta   meta        `json:"meta"`
		Errors []ErrorType `json:"errors"`
	}
)
