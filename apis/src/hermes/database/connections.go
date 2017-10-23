package database

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"hermes/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	"github.com/labstack/echo"
)

// Load enviroment variables
var (
	// Load read database settings
	readDBPort     = utils.GetConfig("HERMES_READDB_PORT", "5432")
	readDBHost     = utils.GetConfig("HERMES_READDB_HOST", "localhost")
	readDBName     = os.Getenv("HERMES_READDB_NAME")
	readDBUser     = os.Getenv("HERMES_READDB_USER")
	readDBPassword = os.Getenv("HERMES_READDB_PASSWORD")
	readDBSslMode  = utils.GetConfig("HERMES_READDB_SSLMODE", "disable")

	ReadDBDriver string

	// Load write database settings
	writeDBPort     = utils.GetConfig("HERMES_WRITEDB_PORT", "5432")
	writeDBHost     = utils.GetConfig("HERMES_WRITEDB_HOST", "localhost")
	writeDBName     = os.Getenv("HERMES_WRITEDB_NAME")
	writeDBUser     = os.Getenv("HERMES_WRITEDB_USER")
	writeDBPassword = os.Getenv("HERMES_WRITEDB_PASSWORD")
	writeDBSslMode  = utils.GetConfig("HERMES_READDB_SSLMODE", "disable")

	WriteDBDriver string
)

// GetReadDB connects to the read database and returns a pointer to the connection
func GetReadDB() (*gorm.DB, error) {
	db, err := getDB(readDBHost, readDBPort, readDBName, readDBUser, readDBPassword, writeDBSslMode)

	if err != nil {
		return nil, err
	}

	ReadDBDriver = reflect.ValueOf(db.DB().Driver()).Type().String()

	return db, nil
}

// GetWriteDB connects to the write database and returns a pointer to the connection
func GetWriteDB() (*gorm.DB, error) {
	db, err := getDB(writeDBHost, writeDBPort, writeDBName, writeDBUser, writeDBPassword, writeDBSslMode)

	if err != nil {
		return nil, err
	}

	WriteDBDriver = reflect.ValueOf(db.DB().Driver()).Type().String()

	return db, nil
}

func getDB(host string, port string, db string, user string, password string, sslMode string) (*gorm.DB, error) {
	credentials := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		host,
		port,
		user,
		db,
		sslMode,
		password)

	connection, err := gorm.Open("postgres", credentials)

	if err != nil {
		log.Println("Failed to connect to database. Error: " + err.Error())

		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return connection, nil
}
