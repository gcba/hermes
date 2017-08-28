package schema

import (
	"net/http"

	"github.com/neelance/graphql-go/errors"
)

func connectionError() error {
	return &StatsError{errors.Errorf("Could not connect to database"), http.StatusInternalServerError}
}

func databaseError() error {
	return &StatsError{errors.Errorf("Could not get value from database"), http.StatusInternalServerError}
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

func badRequestError(message string) error {
	return &StatsError{errors.Errorf(message), http.StatusBadRequest}
}
