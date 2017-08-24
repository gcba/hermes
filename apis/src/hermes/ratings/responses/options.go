package responses

import (
	"net/http"

	"hermes/responses"

	"github.com/labstack/echo"
)

type (
	headers struct {
		ContentType *string `json:"Content-Type"`
		Accept      *string `json:"Accept"`
	}

	Endpoint struct {
		Method  string   `json:"method"`
		Path    string   `json:"path"`
		Headers *headers `json:"headers"`
	}

	Options struct {
		Meta      responses.Status `json:"meta"`
		Endpoints []Endpoint       `json:"endpoints"`
	}
)

var (
	contentType = "application/json; charset=UTF-8"
	accept      = "application/json"

	optionsRoot = Endpoint{
		Method:  "OPTIONS",
		Path:    "/",
		Headers: &headers{}}

	optionsRatings = Endpoint{
		Method:  "OPTIONS",
		Path:    "/ratings",
		Headers: &headers{}}

	postRatings = Endpoint{
		Method: "POST",
		Path:   "/ratings",
		Headers: &headers{
			ContentType: &contentType,
			Accept:      &accept}}

	Endpoints = map[string]Endpoint{
		"OptionsRoot":    optionsRoot,
		"OptionsRatings": optionsRatings,
		"PostRatings":    postRatings}
)

func OptionsResponse(endpoints []Endpoint, context echo.Context) error {
	if !context.Response().Committed {
		response := Options{
			Meta:      responses.Statuses[http.StatusOK],
			Endpoints: endpoints}

		return context.JSON(http.StatusOK, &response)
	}

	return nil
}
