package models

import (
	"ratings/database"
	"strings"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetApp(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	name := uniuri.New()
	key := "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())

	app := App{Name: name, Key: key}
	result := writeDb.Create(&app)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*App); ok {
		var result App

		readDb.First(&result, value.ID)
		require.Equal(t, name, value.Name)
		require.Equal(t, key, value.Key)
	} else {
		t.Fatal("Value is not an App")
	}
}
