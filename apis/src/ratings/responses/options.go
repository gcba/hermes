package responses

type {
	headers struct {
		ContentType string `json:"Content-Type"`
	}

	Endpoint struct {
		Method  string   `json:"method"`
		Path    string   `json:"path"`
		Headers *headers `json:"headers"`
	}

	Options struct {
		Meta      Meta       `json:"meta"`
		Endpoints []Endpoint `json:"endpoints"`
	}
}

const OptionsRoot = Endpoint{
	Method: "OPTIONS",
	Path: "/"
}

const OptionsRatings = Endpoint{
	Method: "OPTIONS",
	Path: "/ratings"
}

const PostRatings = Endpoint{
	Method: "POST",
	Path: "/ratings",
	Headers: &headers{
		ContentType: "application/json; charset=utf-8"
	}
}