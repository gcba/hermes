package parser

import (
	"net/http"

	base "hermes/parser"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
)

type (
	app struct {
		Key     string `json:"key" validate:"required,len=32,alphanum" conform:"trim,lower"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"trim,lower"`
	}

	user struct {
		Name   string `json:"name" validate:"omitempty,gte=3,lte=70" conform:"trim,name"`
		Email  string `json:"email" validate:"omitempty,email,gte=3,lte=100,excludesall= " conform:"trim,email"`
		MiBAID string `json:"mibaId" validate:"omitempty,uuid4,len=36,excludesall= " conform:"trim,lower"`
	}

	platform struct {
		Key     string `json:"key" validate:"required,len=32,alphanum" conform:"trim,lower"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"trim,lower"`
	}

	device struct {
		Name   string  `json:"name" validate:"required,gte=1,lte=30" conform:"trim"`
		Brand  *string `json:"brand" validate:"omitempty,gte=1,lte=30" conform:"trim"`
		Screen screen  `json:"screen" validate:"required"`
	}

	screen struct {
		Width  int  `json:"width" validate:"required,gt=0"`
		Height int  `json:"height" validate:"required,gt=0"`
		PPI    *int `json:"ppi" validate:"omitempty,gt=0"`
	}

	browser struct {
		Name    string `json:"name" validate:"required,gte=1,lte=15" conform:"trim"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"trim,lower"`
	}

	// Request holds the mapped fields from the request's JSON body
	Request struct {
		Rating      int8     `json:"rating" validate:"min=-127,max=127"`
		Description string   `json:"description" validate:"omitempty,gte=3,lte=30" conform:"trim,title"`
		Comment     string   `json:"comment" validate:"omitempty,gte=3,lte=1000" conform:"trim,ucfirst"`
		Range       string   `json:"range" validate:"required,len=32,alphanum" conform:"trim,lower"`
		App         app      `json:"app" validate:"required"`
		Platform    platform `json:"platform" validate:"required"`
		Device      device   `json:"device" validate:"required"`
		User        *user    `json:"user" validate:"omitempty"`
		Browser     *browser `json:"browser" validate:"omitempty"`
	}
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(context echo.Context) (*Request, error) {
	rawRequest, err := base.Parse(newRequest, escape, context)
	request, castOk := rawRequest.(*Request)

	if err != nil {
		return nil, err
	}

	if !castOk {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return request, nil
}

func newRequest() interface{} {
	return new(Request)
}

func escape(request interface{}) error {
	sanitizer := bluemonday.StrictPolicy()

	if data, ok := request.(*Request); ok {
		data.Comment = sanitizer.Sanitize(data.Comment)
		data.Description = sanitizer.Sanitize(data.Description)

		return nil
	}

	return echo.NewHTTPError(http.StatusInternalServerError)
}

func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterStructValidation(UserCustomValidator, user{})
}

func UserCustomValidator(sl validator.StructLevel) {
	item := sl.Current().Interface().(user)

	if len(item.Email) == 0 && len(item.MiBAID) == 0 {
		sl.ReportError(item.Email, "Email", "email", "email/mibaid", "")
		sl.ReportError(item.MiBAID, "MiBAID", "baid", "email/mibaid", "")
	}
}
