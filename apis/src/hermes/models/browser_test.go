package models

import (
	"hermes/database"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateBrowser(t *testing.T) {
	db := database.GetWriteDB()
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
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	name := uniuri.New()
	name = name[3:len(name)]

	browser := Browser{Name: name}
	record := writeDb.Create(&browser)

	if value, ok := record.Value.(*Browser); ok {
		var result Browser

		readDb.First(&result, value.ID)
		require.Equal(t, name, result.Name)
	} else {
		t.Fatal("Value is not a Browser")
	}
}
