package schema

import (
	"hermes/database"
	"strings"

	"github.com/iancoleman/strcase"
)

func fieldExists(field string, fields []string) bool {
	for _, item := range fields {
		if strings.ToUpper(item) == strings.ToUpper(strcase.ToCamel(field)) {
			return true
		}
	}

	return false
}

func isPostgres() bool {
	if database.ReadDBDriver == "*pq.Driver" {
		return true
	}

	return false
}
