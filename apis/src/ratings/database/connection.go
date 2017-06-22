package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite driver
)

// Load environment
var env = os.Getenv("API_RATINGS_ENV")

// Load read database settings
var readDBName = os.Getenv("API_RATINGS_READDB_NAME")
var readDBHost = os.Getenv("API_RATINGS_READDB_HOST")
var readDBUser = os.Getenv("API_RATINGS_READDB_USER")
var readDBPassword = os.Getenv("API_RATINGS_READDB_PASSWORD")

// Load write database settings
var writeDBName = os.Getenv("API_RATINGS_WRITEDB_NAME")
var writeDBHost = os.Getenv("API_RATINGS_WRITEDB_HOST")
var writeDBUser = os.Getenv("API_RATINGS_WRITEDB_USER")
var writeDBPassword = os.Getenv("API_RATINGS_WRITEDB_PASSWORD")

func connectDB(driver string, credentials string) *gorm.DB {
	db, err := gorm.Open(driver, credentials)

	if err != nil {
		panic("Failed to connect to " + driver + " database. Error: " + err.Error())
	}

	return db
}

// GetReadDB connects to the read database and returns a pointer to the connection
func GetReadDB() *gorm.DB {
	if env != "PRODUCTION" {
		credentials := fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
			readDBHost,
			readDBUser,
			readDBName,
			readDBPassword)

		return connectDB("postgres", credentials)
	}

	credentials := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
		readDBHost,
		readDBUser,
		readDBName,
		readDBPassword)

	return connectDB("postgres", credentials)
}

// GetWriteDB connects to the write database and returns a pointer to the connection
func GetWriteDB() *gorm.DB {
	if env != "PRODUCTION" {
		credentials := fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
			readDBHost,
			readDBUser,
			readDBName,
			readDBPassword)

		return connectDB("postgres", credentials)
	}

	credentials := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
		writeDBHost,
		writeDBUser,
		writeDBName,
		writeDBPassword)

	return connectDB("postgres", credentials)
}
