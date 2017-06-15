package responses

type {
	Headers struct {
		ContentType string `json:"Content-Type"`
	}

	Method struct {
		Verb string `json:"verb"`
		Endpoint string `json:"endpoint"`
		Headers Headers `json:"headers"`
	}

	Options struct {
		Meta Meta `json:"meta"`
		Methods []Method `json:"methods"`
	}
}