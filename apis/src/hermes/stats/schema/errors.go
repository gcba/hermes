package schema

import (
	"net/http"

	"github.com/neelance/graphql-go/errors"
)

func connectionError() error {
	return internalServerError("Could not connect to database")
}

func databaseError() error {
	return internalServerError("Could not get value from database")
}

func queryError(err error) error {
	return &StatsError{
		errors.Errorf("Error getting value from database: %v", err.Error()),
		http.StatusInternalServerError}
}

func invalidTableError(table string) error {
	return &StatsError{errors.Errorf("%s is not a valid table", table), http.StatusBadRequest}
}

func invalidFieldError(field string) error {
	return &StatsError{errors.Errorf("%s is not a valid field", field), http.StatusBadRequest}
}

func invalidValueError() error {
	return &StatsError{errors.Errorf("Invalid value provided"), http.StatusBadRequest}
}

func noValueError() error {
	return &StatsError{errors.Errorf("No value provided"), http.StatusBadRequest}
}

func internalServerError(message string) error {
	return &StatsError{errors.Errorf(message), http.StatusInternalServerError}
}

func badRequestError(message string) error {
	return &StatsError{errors.Errorf(message), http.StatusBadRequest}
}
