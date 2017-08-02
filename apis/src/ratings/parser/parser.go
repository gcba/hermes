package parser

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/microcosm-cc/bluemonday"
)

type (
	app struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"trim,lower"`
		Version string `json:"version" validate:"required,gte=1,lte=15,excludesall= " conform:"trim,lower"`
	}

	user struct {
		Name   string `json:"name" validate:"omitempty,gte=3,lte=70" conform:"trim,name"`
		Email  string `json:"email" validate:"omitempty,email,gte=3,lte=100,excludesall= " conform:"trim,email"`
		MiBAID string `json:"mibaId" validate:"omitempty,alphanum,len=36,excludesall= " conform:"trim,lower"`
	}

	platform struct {
		Key     string `json:"key" validate:"required,len=32,alphanum,excludesall= " conform:"trim,lower"`
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
		Range       string   `json:"range" validate:"required,len=32,alphanum,excludesall= " conform:"trim,lower"`
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
		return request, err
	}

	if err := validate(request, context); err != nil {
		return request, err
	}

	return request, nil
}

func bind(request *Request, context echo.Context) error {
	if err := context.Bind(request); err != nil {
		errorDescription := err.Error()
		errorMessage := fmt.Sprintf("Error parsing request: %s", errorDescription)
		errorCode := http.StatusBadRequest

		if httpError, ok := err.(*echo.HTTPError); ok {
			if value, isString := httpError.Message.(string); isString {
				errorMessage = value
				errorCode = httpError.Code
			}
		}

		context.Logger().Error("Error binding request: ", errorDescription)

		return echo.NewHTTPError(errorCode, []string{errorMessage})
	}

	return nil
}

func validate(request *Request, context echo.Context) error {
	if errs := context.Validate(request); errs != nil {
		var errorList []string
		var errorMessage = "Error validating request: "

		if _, ok := errs.(*validator.InvalidValidationError); ok {
			context.Logger().Error(errorMessage, errs.Error())

			return echo.NewHTTPError(http.StatusUnprocessableEntity, []string{errs.Error()})
		}

		for _, err := range errs.(validator.ValidationErrors) {
			errorDescription := err.(error).Error()
			errorList = append(errorList, errorDescription)

			context.Logger().Error(errorMessage, errorDescription)
		}

		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorList)
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
