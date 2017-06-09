package controller

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
)

// Load environment
const env string = os.Getenv("API_RATINGS_ENV")

// Load database settings
const readDatabaseURL string = os.Getenv("API_RATINGS_READ_DB_URL")
const readDatabasePort string = os.Getenv("API_RATINGS_READ_DB_PORT")
const writeDatabaseURL string = os.Getenv("API_RATINGS_WRITE_DB_URL")
const writeDatabasePort string = os.Getenv("API_RATINGS_WRITE_DB_PORT")

func connectDB(name string, url string, port string) *gorm.DB {
	db, err := gorm.Open(url, port)
	if err != nil {
		panic("Failed to connect to " + name)
	}
	defer db.Close()

	return db
}

func getReadDB() *gorm.DB {
	if env != "PRODUCTION" {
		return connectDB("Read Database", "../"+readDatabaseURL, readDatabasePort)
	}

	return connectDB("Read Database", "../"+readDatabaseURL, readDatabasePort)
}

func getWriteDB() *gorm.DB {
	if env != "PRODUCTION" {
		return connectDB("Write Database", "../"+writeDatabaseURL, writeDatabasePort)
	}

	return connectDB("Write Database", "../"+writeDatabaseURL, writeDatabasePort)
}
