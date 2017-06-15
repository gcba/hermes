package responses

type {
	Headers struct {
		ContentType string `json:"Content-Type"`
	}

	Method struct {
		Verb string `json:"verb"`
		Headers Headers `json:"headers"`
	}

	Options struct {
		Meta Meta `json:"meta"`
		Methods []Method `json:"methods"`
	}
}