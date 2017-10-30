package models

import (
	"strings"
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetApp(t *testing.T) {
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

	app := App{
		Name: uniuri.New(),
		Key:  "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())}

	result := writeDb.Create(&app)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*App); ok {
		var result App

		readDb.First(&result, value.ID)
		require.Equal(t, app.Name, value.Name)
		require.Equal(t, app.Key, value.Key)
	} else {
		t.Fatal("Value is not an App")
	}
}
