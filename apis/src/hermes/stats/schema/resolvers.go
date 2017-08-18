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
