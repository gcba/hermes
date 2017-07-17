package parser

import (
	"fmt"
	"net/http"

	"ratings/responses"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/microcosm-cc/bluemonday"
)

type (
	app struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"lower"`
	}

	user struct {
		Name   string `json:"name" validate:"omitempty,gte=3,lte=70" conform:"name"`
		Email  string `json:"email" validate:"omitempty,email,gte=3,lte=100,excludesall= " conform:"email"`
		MiBAID string `json:"mibaId" validate:"omitempty,alphanum,gte=1,excludesall= " conform:"lower"`
	}

	platform struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"lower"`
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
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"lower"`
	}

	// Request holds the mapped fields from the request's JSON body
	Request struct {
		Rating      int8     `json:"rating" validate:"min=-127,max=127"`
		Description string   `json:"description" validate:"omitempty,gte=3,lte=30" conform:"trim,title"`
		Comment     string   `json:"comment" validate:"omitempty,gte=3,lte=1000" conform:"trim,ucfirst"`
		Range       string   `json:"range" validate:"required,len=32,alphanum,excludesall= " conform:"lower"`
		App         app      `json:"app" validate:"required"`
		Platform    platform `json:"platform" validate:"required"`
		Device      device   `json:"device" validate:"required"`
		User        *user    `json:"user" validate:"omitempty"`
		Browser     *browser `json:"browser" validate:"omitempty"`
	}
)

// Parse parses, scrubs and escapes a request's JSON body and maps it to a struct
func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	conform.Strings(request)
	escape(request)

	if err := bind(request, context); err != nil {
		context.Logger().Error("Error binding request: " + err.Error())

		return request, err
	}

	if err := validate(request, context); err != nil {
		context.Logger().Error("Error validating request: " + err.Error())

		return request, err
	}

	return request, nil
}

func bind(request *Request, context echo.Context) error {
	if err := context.Bind(request); err != nil {
		errorMessage := fmt.Sprintf("Error parsing request: %s", err.Error())
		errorCode := http.StatusBadRequest

		if httpError, ok := err.(*echo.HTTPError); ok {
			if value, isString := httpError.Message.(string); isString {
				errorMessage = value
				errorCode = httpError.Code
			}
		}

		context.Logger().Error(errorMessage)

		return responses.ErrorResponse(errorCode, errorMessage, context)
	}

	return nil
}

func validate(request *Request, context echo.Context) error {
	if errs := context.Validate(request); errs != nil {
		var errorList []string

		if _, ok := errs.(*validator.InvalidValidationError); ok {
			return responses.ErrorResponse(http.StatusUnprocessableEntity, errs.Error(), context)
		}

		for _, err := range errs.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error validating request: %s", err.(error).Error())
			errorList = append(errorList, errorMessage)

			context.Logger().Error(errorMessage)
		}

		return responses.ErrorsResponse(http.StatusUnprocessableEntity, errorList, context)
	}

	return nil
}

func escape(request *Request) {
	sanitizer := bluemonday.StrictPolicy()

	request.Comment = sanitizer.Sanitize(request.Comment)
	request.Description = sanitizer.Sanitize(request.Description)
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
