package models

import (
	"hermes/database"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateDevice(t *testing.T) {
	db := database.GetWriteDB()
	defer db.Close()

	name := uniuri.New()

	device := Device{
		Name:         name,
		ScreenWidth:  720,
		ScreenHeight: 1280,
		PPI:          320,
		BrandID:      1,
		PlatformID:   2}

	result := db.Create(&device)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Device); ok {
		require.Equal(t, name, value.Name)
		require.Equal(t, device.ScreenHeight, value.ScreenHeight)
		require.Equal(t, device.ScreenWidth, value.ScreenWidth)
		require.Equal(t, device.PPI, value.PPI)
		require.Equal(t, device.BrandID, value.BrandID)
		require.Equal(t, device.PlatformID, value.PlatformID)
	} else {
		t.Fatal("Value is not a Device")
	}
}

func TestGetDevice(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	name := uniuri.New()

	device := Device{
		Name:         name,
		ScreenWidth:  960,
		ScreenHeight: 1600,
		PPI:          430,
		BrandID:      2,
		PlatformID:   2}

	record := writeDb.Create(&device)

	if value, ok := record.Value.(*Device); ok {
		var result Device

		readDb.First(&result, value.ID)

		require.Equal(t, name, value.Name)
		require.Equal(t, device.ScreenHeight, value.ScreenHeight)
		require.Equal(t, device.ScreenWidth, value.ScreenWidth)
		require.Equal(t, device.PPI, value.PPI)
		require.Equal(t, device.BrandID, value.BrandID)
		require.Equal(t, device.PlatformID, value.PlatformID)
	} else {
		t.Fatal("Value is not a Device")
	}
}
