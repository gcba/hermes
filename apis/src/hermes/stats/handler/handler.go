package handler

import (
	"io/ioutil"
	"net/http"
	"os"
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
		db *gorm.DB
	}

	CountResolver struct {
		db *gorm.DB
	}

	AverageResolver struct {
		db *gorm.DB
	}
)

func (cr *CountResolver) Ratings(context echo.Context, args struct{ Field field }) (int, error) {
	return 0, nil
}

func (ar *AverageResolver) Ratings(context echo.Context, args struct{ Field field }) (float64, error) {
	return 0, nil
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
