package handler

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"hermes/database"
	"hermes/stats/responses"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/jinzhu/gorm"
	"github.com/neelance/graphql-go"
)

// SCHEMA --- Extract this into another package

var (
	schema *graphql.Schema
)

type (
	field struct {
		Name     string
		Operator *string
		Int      *int
		Float    *float64
		String   *string
		Bool     *bool
		Next     *operation
	}

	operation struct {
		Condition string
		Field     *field
	}

	Resolver struct {
		db: *gorm.DB
	}

	CountResolver struct {
		db: *gorm.DB
	}

	AverageResolver struct {
		db: *gorm.DB
	}
)

func (cr *CountResolver) Ratings(args struct{ Field field }) int {
	return 0
}

func (ar *AverageResolver) Ratings(args struct{ Field field }) float {
	return 0
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{db: db}
}

func ParseSchema() {
	var rawSchema []byte
	var err error
	var db = database.GetReadDB()

	defer db.Close()

	rawSchema, err = ioutil.ReadFile("../schema/schema.graphql")

	if err != nil {
		panic(err)
	}

	schema, err = graphql.ParseSchema(string(rawSchema), NewResolver(db))

	if err != nil {
		panic(err)
	}
}

// END SCHEMA --- Extract this into another package

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func badRequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var message string

		if !hasAcceptHeader(context) {
			message = "Accept header is missing"

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		if context.Request().Method == echo.OPTIONS && hasContentTypeHeader(context) {
			message = "OPTIONS requests must have no body"

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		if context.Request().Method == echo.POST && !hasContentTypeHeader(context) {
			message = "Content-Type header is missing"

			return echo.NewHTTPError(http.StatusBadRequest, []string{message})
		}

		return next(context)
	}
}

func notAcceptableMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var message string

		if !isValidAcceptHeader(context) {
			message = "JSON responses must be accepted"

			return echo.NewHTTPError(http.StatusNotAcceptable, []string{message})
		}

		return next(context)
	}
}

func notImplementedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if context.Request().Method != echo.POST && context.Request().Method != echo.OPTIONS {
			return echo.NewHTTPError(http.StatusNotImplemented)
		}

		return next(context)
	}
}

func corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		context.Response().Header().Set("Access-Control-Allow-Methods", "POST")
		context.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

		return next(context)
	}
}

func hasAcceptHeader(context echo.Context) bool {
	if header := context.Request().Header.Get("Accept"); strings.TrimSpace(header) != "" {
		return true
	}

	return false
}

func isValidAcceptHeader(context echo.Context) bool {
	accepted := "application/json"

	if header := context.Request().Header.Get("Accept"); strings.Contains(strings.ToLower(header), accepted) || header == "*/*" {
		return true
	}

	return false
}

func hasContentTypeHeader(context echo.Context) bool {
	if header := context.Request().Header.Get("Content-Type"); strings.TrimSpace(header) != "" {
		return true
	}

	return false
}

func Handler(port int, handlers map[string]echo.HandlerFunc) http.Handler {
	e := echo.New()
	validate := validator.New()
	env := os.Getenv("HERMES_STATS_ENV")

	if env == "DEV" {
		e.Logger.SetLevel(log.DEBUG)
		e.Debug = true
	} else {
		e.Pre(middleware.HTTPSRedirect())
		e.Logger.SetLevel(log.ERROR)
	}

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("20K"))
	e.Use(notImplementedMiddleware)
	e.Use(notAcceptableMiddleware)
	e.Use(badRequestMiddleware)
	e.Use(corsMiddleware)

	e.POST("/stats", handlers["PostStats"])

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validate}
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}
