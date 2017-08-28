package models

import (
	"hermes/database"
)

func isPostgres() bool {
	if database.ReadDBDriver == "*pq.Driver" {
		return true
	}

	return false
}
