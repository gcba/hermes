package schema

import (
	"context"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	field struct {
		Name     string
		Operator *string
		Value    Value
		Next     *operation
	}

	operation struct {
		Condition string
		Field     *field
	}

	Resolver struct {
		db *gorm.DB
	}
)

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func (r *Resolver) Count(context context.Context, args struct{ Field *field }) (int32, error) {
	// TODO: Implement

	return 0, nil
}

func (r *Resolver) Average(context context.Context, args struct{ Field *field }) (float64, error) {
	// TODO: Implement

	return 0, nil
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{db: db}
}
