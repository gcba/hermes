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

const (
	OptionsRoot = Endpoint{
		Method: "OPTIONS",
		Path: "/"
	}

	OptionsRatings = Endpoint{
		Method: "OPTIONS",
		Path: "/ratings"
	}

	PostRatings = Endpoint{
		Method: "POST",
		Path: "/ratings",
		Headers: &headers{
			ContentType: "application/json; charset=utf-8"
		}
	}
)

func OptionsResponse(status int) *Options {

}