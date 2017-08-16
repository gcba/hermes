package models

import (
	"hermes/database"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateBrand(t *testing.T) {
	db := database.GetWriteDB()
	defer db.Close()

	name := uniuri.New()
	brand := Brand{Name: name}
	result := db.Create(&brand)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Brand); ok {
		require.Equal(t, name, value.Name)
	} else {
		t.Fatal("Value is not a Brand")
	}
}

func TestGetBrand(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	name := uniuri.New()
	brand := Brand{Name: name}
	record := writeDb.Create(&brand)

	if value, ok := record.Value.(*Brand); ok {
		var result Brand

		readDb.First(&result, value.ID)
		require.Equal(t, name, result.Name)
	} else {
		t.Fatal("Value is not a Brand")
	}
}
