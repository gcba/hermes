package models

import (
	"hermes/database"
	"strings"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetPlatform(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	name := uniuri.New()
	name = name[3:len(name)]

	platform := Platform{
		Name: name,
		Key:  "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())}

	result := writeDb.Create(&platform)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Platform); ok {
		var result Platform

		readDb.First(&result, value.ID)
		require.Equal(t, platform.Name, value.Name)
		require.Equal(t, platform.Key, value.Key)
	} else {
		t.Fatal("Value is not a Platform")
	}
}
