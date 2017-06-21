package parser

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"ratings/responses"
)

type (
	app struct {
		Key     string `json:"key" validate:"required,len=32,alphanum"` // The app key
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15"`
	}

	user struct {
		Name   string `json:"name" validate:"required,alphanum,gte=3,slte=70"`
		Email  string `json:"email" validate:"required,email,gte=3,lte=100"`
		MiBAID string `json:"mibaId" validate:"required,alphanum,gte=1"`
	}

	platform struct {
		Key     string `json:"key" validate:"required,len=32,alphanum"` // The platform key
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15"`
	}

	device struct {
		Name   string `json:"name" validate:"required,alphanum,gte=1,lte=30"`
		Brand  string `json:"brand" validate:"required,alphanum,gte=1,lte=30"`
		Screen screen `json:"screen" validate:"required"`
	}

	screen struct {
		Width  int `json:"width" validate:"required,numeric,gt=0"`
		Height int `json:"height" validate:"required,numeric,gt=0"`
		PPI    int `json:"ppi" validate:"required,numeric,gt=0"`
	}

	browser struct {
		Name    string `json:"name" validate:"required,alphanum,gte=1,lte=15"`
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15"`
	}

	// Request holds the mapped fields from the request's JSON body
	Request struct {
		Rating      uint8    `json:"rating" validate:"required,numeric,min=-127,max=127"`
		Description string   `json:"description" validate:"alphanum,gte=3,lte=30,omitempty"`
		Comment     string   `json:"comment" validate:"alphanum,gte=3,lte=1000,omitempty"`
		Range       string   `json:"range" validate:"required,len=32,alphanum"` // The range key
		App         app      `json:"app" validate:"required"`
		Platform    platform `json:"platform" validate:"required"`
		User        *user    `json:"user" validate:"omitempty"`
		Device      *device  `json:"device" validate:"omitempty"`
		Browser     *browser `json:"browser" validate:"omitempty"`
	}
)

// Parse parses a request's JSON body and maps it to a struct
func Parse(context echo.Context) *Request {
	request := new(Request)

	if err := context.Bind(request); err != nil {
		errorMessage := fmt.Sprintf("Error parsing request: %s", err.Error())

		responses.ErrorResponse(http.StatusBadRequest, errorMessage, context)

		return request
	}

	if err := context.Validate(request); err != nil {
		errorMessage := fmt.Sprintf("Error validating request: %s", err.Error())

		responses.ErrorResponse(http.StatusUnprocessableEntity, errorMessage, context)

		return request
	}

	return request
}
