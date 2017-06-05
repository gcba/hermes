package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("../../../../admin/database", "database.sqlite")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	return db
}
