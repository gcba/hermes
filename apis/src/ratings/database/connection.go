package database

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
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
	readDBPort     = os.Getenv("API_RATINGS_READDB_PORT")
	readDBHost     = os.Getenv("API_RATINGS_READDB_HOST")
	readDBName     = os.Getenv("API_RATINGS_READDB_NAME")
	readDBUser     = os.Getenv("API_RATINGS_READDB_USER")
	readDBPassword = os.Getenv("API_RATINGS_READDB_PASSWORD")

	// Load write database settings
	writeDBPort     = os.Getenv("API_RATINGS_WRITEDB_PORT")
	writeDBHost     = os.Getenv("API_RATINGS_WRITEDB_HOST")
	writeDBName     = os.Getenv("API_RATINGS_WRITEDB_NAME")
	writeDBUser     = os.Getenv("API_RATINGS_WRITEDB_USER")
	writeDBPassword = os.Getenv("API_RATINGS_WRITEDB_PASSWORD")
)

// GetReadDB connects to the read database and returns a pointer to the connection
func GetReadDB() *gorm.DB {
	return getDB(readDBHost, readDBPort, readDBName, readDBUser, readDBPassword)
}

// GetWriteDB connects to the write database and returns a pointer to the connection
func GetWriteDB() *gorm.DB {
	return getDB(writeDBHost, writeDBPort, writeDBName, writeDBUser, writeDBPassword)
}

func getDB(host string, port string, db string, user string, password string) *gorm.DB {
	if len(host) == 0 {
		host = "localhost"
	}

	if len(port) == 0 {
		port = "5432"
	}

	if env != "PRODUCTION" {
		credentials := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
			host,
			port,
			user,
			db,
			password)

		return connectDB("postgres", credentials)
	}

	credentials := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
		host,
		port,
		user,
		db,
		password)

	return connectDB("postgres", credentials)
}

func connectDB(driver string, credentials string) *gorm.DB {
	if envErr != nil {
		log.Fatal("Error loading .env file: ", envErr.Error())
	}

	db, dbErr := gorm.Open(driver, credentials)

	if dbErr != nil {
		log.Println("Failed to connect to " + driver + " database. Error: " + dbErr.Error())
	}

	return db
}
