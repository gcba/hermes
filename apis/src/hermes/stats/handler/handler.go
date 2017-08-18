package handler

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"

	"hermes/database"
	"hermes/middlewares"
	"hermes/stats/responses"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/neelance/graphql-go"
)

// SCHEMA --- Extract this into another package
const DB = iota

var (
	Schema *graphql.Schema
)

type (
	field struct {
		Name     string
		Operator *string
		Int      *int32
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
		db *gorm.DB
	}

	CountResolver struct {
		db *gorm.DB
	}

	AverageResolver struct {
		db *gorm.DB
	}
)

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func (r *Resolver) Count(context context.Context) (*CountResolver, error) {
	if db, ok := context.Value(DB).(*gorm.DB); ok {
		// TODO: Implement

		return &CountResolver{db: db}, nil
	}

	return nil, errorResponse()
}

func (r *Resolver) Average(context context.Context) (*AverageResolver, error) {
	if db, ok := context.Value(DB).(*gorm.DB); ok {
		// TODO: Implement

		return &AverageResolver{db: db}, nil
	}

	return nil, errorResponse()
}

func (cr *CountResolver) Ratings(context context.Context, args struct{ Field *field }) (int32, error) {
	// TODO: Implement

	return 0, nil
}

func (ar *AverageResolver) Ratings(context context.Context, args struct{ Field *field }) (float64, error) {
	// TODO: Implement

	return 0, nil
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{db: db}
}

func ParseSchema() {
	var rawSchema []byte
	var err error

	db := database.GetReadDB()

	defer db.Close()

	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not load GraphQL schema")
	}

	rawSchema, err = ioutil.ReadFile(path.Join(path.Dir(filename), "../schema/schema.graphql"))

	if err != nil {
		panic(err)
	}

	Schema, err = graphql.ParseSchema(string(rawSchema), NewResolver(db))

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

func init() {
	ParseSchema()
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
	e.Use(middlewares.NotImplementedMiddleware)
	e.Use(middlewares.NotAcceptableMiddleware)
	e.Use(middlewares.BadRequestMiddleware)
	e.Use(middlewares.UnsupportedMediaTypeMiddleware)
	e.Use(middlewares.CorsMiddleware)

	e.POST("/stats", handlers["PostStats"])

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validate}
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}
