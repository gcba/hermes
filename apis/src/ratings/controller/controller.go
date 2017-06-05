package controller

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
)

func connectDB() *gorm.DB {
	env := os.Getenv("API_RATINGS_ENV")

	if env != "PRODUCTION" {
		db, err := gorm.Open("../../../../admin/database", "database.sqlite")
		if err != nil {
			panic("Failed to connect to database")
		}
		defer db.Close()

		return db
	}

	db, err := gorm.Open("../../../../admin/database", "database.sqlite")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	return db
}
