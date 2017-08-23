package schema

import (
	"context"
	"net/http"

	"github.com/deckarep/golang-set"
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
		Field     field
	}

	Resolver struct{}
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

func (f *field) walk(callback func(*field)) {
	if f.Next != nil {
		f.Next.Field.walk(callback)
	}

	go callback(f)
}

func (f *field) toSet() mapset.Set {
	return mapset.NewSet()
}

func getTableName() {

}

func getFieldName() {

}
