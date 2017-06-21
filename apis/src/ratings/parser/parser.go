package parser

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"ratings/responses"

	"github.com/labstack/echo"
	"github.com/leebenson/conform"
)

type (
	app struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15,excludesall= " conform:"lower"`
	}

	user struct {
		Name   string `json:"name" validate:"required,alphanum,gte=3,slte=70" conform:"name"`
		Email  string `json:"email" validate:"required,email,gte=3,lte=100,excludesall= " conform:"email"`
		MiBAID string `json:"mibaId" validate:"required,alphanum,gte=1,excludesall= " conform:"lower"`
	}

	platform struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15,excludesall= " conform:"lower"`
	}

	device struct {
		Name   string `json:"name" validate:"required,alphanum,gte=1,lte=30" conform:"trim"`
		Brand  string `json:"brand" validate:"required,alphanum,gte=1,lte=30" conform:"trim"`
		Screen screen `json:"screen" validate:"required"`
	}

	screen struct {
		Width  int `json:"width" validate:"required,numeric,gt=0"`
		Height int `json:"height" validate:"required,numeric,gt=0"`
		PPI    int `json:"ppi" validate:"required,numeric,gt=0"`
	}

	browser struct {
		Name    string `json:"name" validate:"required,alphanum,gte=1,lte=15" conform:"trim"`
		Version string `json:"version" validate:"required,alphanum,gte=1,lte=15,excludesall= " conform:"lower"`
	}

	// Request holds the mapped fields from the request's JSON body
	Request struct {
		Rating      uint8    `json:"rating" validate:"required,numeric,min=-127,max=127"`
		Description string   `json:"description" validate:"alphanum,gte=3,lte=30,omitempty" conform:"trim"`
		Comment     string   `json:"comment" validate:"gte=3,lte=1000,omitempty" conform:"trim,ucfirst"`
		Range       string   `json:"range" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		App         app      `json:"app" validate:"required"`
		Platform    platform `json:"platform" validate:"required"`
		User        *user    `json:"user" validate:"omitempty"`
		Device      *device  `json:"device" validate:"omitempty"`
		Browser     *browser `json:"browser" validate:"omitempty"`
	}
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	if err := bind(request, context); err != nil {
		return request, err
	}

	if err := validate(request, context); err != nil {
		return request, err
	}

	conform.Strings(request)

	if err := escapeComment(request, context); err != nil {
		return request, err
	}

	return request, nil
}

func bind(request *Request, context echo.Context) error {
	if err := context.Bind(request); err != nil {
		errorMessage := fmt.Sprintf("Error parsing request: %s", err.Error())

		return responses.ErrorResponse(http.StatusBadRequest, errorMessage, context)
	}

	return nil
}

func validate(request *Request, context echo.Context) error {
	if err := context.Validate(request); err != nil {
		errorMessage := fmt.Sprintf("Error validating request: %s", err.Error())

		return responses.ErrorResponse(http.StatusUnprocessableEntity, errorMessage, context)
	}

	return nil
}

func escapeComment(request *Request, context echo.Context) error {
	comment := bytes.NewBufferString("")
	commentTemplate, commentTemplateErr := template.New("Comment").Parse(`{{define "C"}{{.}}{{end}}`)

	if commentTemplateErr != nil {
		errorMessage := fmt.Sprintf("Error escaping comment: %s", commentTemplateErr.Error())

		return responses.ErrorResponse(http.StatusInternalServerError, errorMessage, context)
	}

	commentErr := commentTemplate.ExecuteTemplate(comment, "C", request.Comment)

	if commentErr != nil {
		errorMessage := fmt.Sprintf("Error escaping comment: %s", commentErr.Error())

		return responses.ErrorResponse(http.StatusInternalServerError, errorMessage, context)
	}

	request.Comment = comment.String()

	return nil
}
