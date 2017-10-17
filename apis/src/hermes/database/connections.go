package database

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

// Load enviroment variables
var (
	// Load read database settings
	readDBPort     = getConfig("HERMES_READDB_PORT", "5432")
	readDBHost     = getConfig("HERMES_READDB_HOST", "localhost")
	readDBName     = os.Getenv("HERMES_READDB_NAME")
	readDBUser     = os.Getenv("HERMES_READDB_USER")
	readDBPassword = os.Getenv("HERMES_READDB_PASSWORD")

	ReadDBDriver string

	// Load write database settings
	writeDBPort     = getConfig("HERMES_WRITEDB_PORT", "5432")
	writeDBHost     = getConfig("HERMES_WRITEDB_HOST", "localhost")
	writeDBName     = os.Getenv("HERMES_WRITEDB_NAME")
	writeDBUser     = os.Getenv("HERMES_WRITEDB_USER")
	writeDBPassword = os.Getenv("HERMES_WRITEDB_PASSWORD")

	WriteDBDriver string
)

// GetReadDB connects to the read database and returns a pointer to the connection
func GetReadDB() *gorm.DB {
	db := getDB(readDBHost, readDBPort, readDBName, readDBUser, readDBPassword)

	ReadDBDriver = reflect.ValueOf(db.DB().Driver()).Type().String()

	return db
}

// GetWriteDB connects to the write database and returns a pointer to the connection
func GetWriteDB() *gorm.DB {
	db := getDB(writeDBHost, writeDBPort, writeDBName, writeDBUser, writeDBPassword)

	WriteDBDriver = reflect.ValueOf(db.DB().Driver()).Type().String()

	return db
}

func getDB(host string, port string, db string, user string, password string) *gorm.DB {
	credentials := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", // TODO: Handle sslmode
		host,
		port,
		user,
		db,
		password)

	connection, err := gorm.Open("postgres", credentials)

	if err != nil {
		log.Println("Failed to connect to database. Error: " + err.Error())
	}

	return connection
}

func getConfig(env string, fallback string) string {
	if result := os.Getenv(env); len(result) > 0 {
		return result
	}

	return fallback
}
