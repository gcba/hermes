package database

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite driver
	"github.com/joho/godotenv"
)

// Load enviroment variables
var (
	rootPath, _ = filepath.Abs(".")
	envPath     = path.Join(rootPath, ".env")
	envErr      = godotenv.Load(envPath)

	// Load environment
	env = os.Getenv("API_RATINGS_ENV")

	// Load read database settings
	readDBName     = os.Getenv("API_RATINGS_READDB_NAME")
	readDBHost     = os.Getenv("API_RATINGS_READDB_HOST")
	readDBUser     = os.Getenv("API_RATINGS_READDB_USER")
	readDBPassword = os.Getenv("API_RATINGS_READDB_PASSWORD")

	// Load write database settings
	writeDBName     = os.Getenv("API_RATINGS_WRITEDB_NAME")
	writeDBHost     = os.Getenv("API_RATINGS_WRITEDB_HOST")
	writeDBUser     = os.Getenv("API_RATINGS_WRITEDB_USER")
	writeDBPassword = os.Getenv("API_RATINGS_WRITEDB_PASSWORD")
)

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

func connectDB(driver string, credentials string) *gorm.DB {
	if envErr != nil {
		log.Fatal("Error loading .env file: ", envErr.Error())
	}

	db, dbErr := gorm.Open(driver, credentials)

	if dbErr != nil {
		log.Fatal("Failed to connect to " + driver + " database. Error: " + dbErr.Error())
	}

	return db
}
