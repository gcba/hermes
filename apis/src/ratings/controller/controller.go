package controller

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite driver
)

// Load environment
const env string = os.Getenv("API_RATINGS_ENV")

// Load database settings
const readDBName string = os.Getenv("API_RATINGS_READDB_NAME")
const readDBHost string = os.Getenv("API_RATINGS_READDB_HOST")
const readDBUser string = os.Getenv("API_RATINGS_READDB_USER")
const readDBPassword string = os.Getenv("API_RATINGS_READDB_PASSWORD")

const writeDBName string = os.Getenv("API_RATINGS_WRITEDB_NAME")
const writeDBHost string = os.Getenv("API_RATINGS_WRITEDB_HOST")
const writeDBUser string = os.Getenv("API_RATINGS_WRITEDB_USER")
const writeDBPassword string = os.Getenv("API_RATINGS_WRITEDB_PASSWORD")

func connectDB(driver string, credentials string) *gorm.DB {
	db, err := gorm.Open(driver, credentials)
	if err != nil {
		panic("Failed to connect to " + driver + " database")
	}
	defer db.Close()

	return db
}

// GetReadDB connects to the read database and return a pointer to the connection
func GetReadDB() *gorm.DB {
	if env != "PRODUCTION" {
		return connectDB("sqlite3", readDBHost+'/'+readDBName)
	}

	credentials := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		readDBHost,
		readDBUser,
		readDBName,
		readDBPassword
		)

	return connectDB("postgres", credentials)
}

// GetWriteDB connects to the write database and return a pointer to the connection
func GetWriteDB() *gorm.DB {
	if env != "PRODUCTION" {
		return connectDB("sqlite3", writeDBHost+'/'+writeDBName)
	}

	credentials := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		writeDBHost,
		writeDBUser,
		writeDBName,
		writeDBPassword
		)

	return connectDB("postgres", credentials)
}
