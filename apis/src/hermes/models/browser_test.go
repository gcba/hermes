package models

import (
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateBrowser(t *testing.T) {
	db, dbError := database.GetWriteDB()

	if dbError != nil {
		t.Fatal("Could not get connect to write database")
	}

	defer db.Close()

	name := uniuri.New()
	name = name[3:len(name)]

	browser := Browser{Name: name}
	result := db.Create(&browser)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Browser); ok {
		require.Equal(t, name, value.Name)
	} else {
		t.Fatal("Value is not a Browser")
	}
}

func TestGetBrowser(t *testing.T) {
	writeDb, writeDbError := database.GetWriteDB()

	if writeDbError != nil {
		t.Fatal("Could not get connect to write database")
	}

	defer writeDb.Close()

	readDb, readDbError := database.GetReadDB()

	if readDbError != nil {
		t.Fatal("Could not get connect to read database")
	}

	defer readDb.Close()

	name := uniuri.New()
	name = name[3:len(name)]

	browser := Browser{Name: name}
	record := writeDb.Create(&browser)

	if value, ok := record.Value.(*Browser); ok {
		require.Equal(t, name, value.Name)
	} else {
		t.Fatal("Value is not a Browser")
	}
}
