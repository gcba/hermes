package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

// Load enviroment variables
var (
	// Load read database settings
	readDBPort     = os.Getenv("HERMES_READDB_PORT")
	readDBHost     = os.Getenv("HERMES_READDB_HOST")
	readDBName     = os.Getenv("HERMES_READDB_NAME")
	readDBUser     = os.Getenv("HERMES_READDB_USER")
	readDBPassword = os.Getenv("HERMES_READDB_PASSWORD")

	// Load write database settings
	writeDBPort     = os.Getenv("HERMES_WRITEDB_PORT")
	writeDBHost     = os.Getenv("HERMES_WRITEDB_HOST")
	writeDBName     = os.Getenv("HERMES_WRITEDB_NAME")
	writeDBUser     = os.Getenv("HERMES_WRITEDB_USER")
	writeDBPassword = os.Getenv("HERMES_WRITEDB_PASSWORD")
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
	db, dbErr := gorm.Open(driver, credentials)

	if dbErr != nil {
		log.Println("Failed to connect to " + driver + " database. Error: " + dbErr.Error())
	}

	return db
}
