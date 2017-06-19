package responses

import (
	"net/http"
)

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
	optionsRoot = Endpoint{
		Method: "OPTIONS",
		Path: "/"
	}

	optionsRatings = Endpoint{
		Method: "OPTIONS",
		Path: "/ratings"
	}

	postRatings = Endpoint{
		Method: "POST",
		Path: "/ratings",
		Headers: &headers{
			ContentType: "application/json; charset=utf-8"
		}
	}

	Endpoints = map[string]Endpoint{
		"optionsRoot": optionsRoot,
		"optionsRatings": optionsRatings,
		"postRatings": postRatings
	}
)

func OptionsResponse(endpoints []Endpoint, context echo.Context) error {
	response := Options{
		Meta: metas[http.StatusOk],
		Endpoints: endpoints
	}

  	return context.JSON(http.StatusOk, &response)
}