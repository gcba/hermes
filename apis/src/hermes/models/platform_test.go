package models

import (
	"strings"
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetPlatform(t *testing.T) {
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

	platform := Platform{
		Name: name,
		Key:  "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())}

	result := writeDb.Create(&platform)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Platform); ok {
		require.Equal(t, platform.Name, value.Name)
		require.Equal(t, platform.Key, value.Key)
	} else {
		t.Fatal("Value is not a Platform")
	}
}
