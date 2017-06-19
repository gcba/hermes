package parser

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	app struct {
		Key     string `json:"key";validate:"required,len=32,alphanum"`
		Version string `json:"version";validate:"required,alphanum,gte=1,lte=15"`
	}

	user struct {
		Name   string `json:"name";validate:"required,alphanum,gte=3,lte=70"`
		Email  string `json:"email";validate:"required,email,gte=3,lte=100"`
		MiBAID string `json:"mibaId";validate:"required,alphanum,gte=1"`
	}

	platform struct {
		Key     string `json:"key";validate:"required,len=32,alphanum"`
		Version string `json:"version";validate:"required,alphanum,gte=1,lte=15"`
	}

	device struct {
		Name   string `json:"name";validate:"required,alphanum,gte=1,lte=30"`
		Brand  string `json:"brand";validate:"required,alphanum,gte=1,lte=30"`
		Screen screen `json:"screen";validate:"required,dive,required"`
	}

	screen struct {
		Width  int `json:"width";validate:"required,numeric,gt=0"`
		Height int `json:"height";validate:"required,numeric,gt=0"`
		PPI    int `json:"ppi";validate:"required,numeric,gt=0"`
	}

	browser struct {
		Name    string `json:"name";validate:"required,alphanum,gte=1,lte=15"`
		Version string `json:"version";validate:"required,alphanum,gte=1,lte=15"`
	}

	Request struct {
		Rating      int8     `json:"rating";validate:"required,numeric,min=-127,max=127"`
		Description string   `json:"description";validate:"alphanum,gte=3,lte=30,omitempty"`
		Comment     string   `json:"comment";validate:"alphanum,gte=3,lte=1000,omitempty"`
		Range       string   `json:"range";validate:"required,len=32,alphanum"` // The range key
		App         app      `json:"app";validate:"required,dive,required"`
		Platform    platform `json:"platform";validate:"required,dive,required"`
		User        *user    `json:"user";validate:"omitempty"`
		Device      *device  `json:"device";validate:"omitempty"`
		Browser     *browser `json:"browser";validate:"omitempty"`
	}

	RequestValidator struct {
		validator *validator.Validate
	}
)

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	if err := context.Bind(request); err != nil {
		fmt.Println("Error parsing request:", err) // TODO: Return error response

		return request, err
	}

	if err = context.Validate(request); err != nil {
		fmt.Println("Error validating request:", err) // TODO: Return error response

		return request, err
	}

	return request, nil
}
