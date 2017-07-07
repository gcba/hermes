package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	headers struct {
		ContentType string `json:"Content-Type"`
	}

	Endpoint struct {
		Method  string   `json:"method"`
		Path    string   `json:"path"`
		Headers *headers `json:"headers"`
	}

	Options struct {
		Meta      meta       `json:"meta"`
		Endpoints []Endpoint `json:"endpoints"`
	}
)

var (
	optionsRoot = Endpoint{
		Method: "OPTIONS",
		Path:   "/",
		Headers: &headers{
			ContentType: "application/json; charset=UTF-8"}}

	optionsRatings = Endpoint{
		Method: "OPTIONS",
		Path:   "/ratings",
		Headers: &headers{
			ContentType: "application/json; charset=UTF-8"}}

	postRatings = Endpoint{
		Method: "POST",
		Path:   "/ratings",
		Headers: &headers{
			ContentType: "application/json; charset=UTF-8"}}

	Endpoints = map[string]Endpoint{
		"OptionsRoot":    optionsRoot,
		"OptionsRatings": optionsRatings,
		"PostRatings":    postRatings}
)

func OptionsResponse(endpoints []Endpoint, context echo.Context) error {
	response := Options{
		Meta:      metas[http.StatusOK],
		Endpoints: endpoints}

	return context.JSON(http.StatusOK, &response)
}
