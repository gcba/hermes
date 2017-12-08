package models

import (
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateDevice(t *testing.T) {
	db, dbError := database.GetWriteDB()

	if dbError != nil {
		t.Fatal("Could not get connect to write database")
	}

	defer db.Close()

	name := uniuri.New()
	ppi := 320
	brandId := uint(1)
	platformId := uint(2)

	device := Device{
		Name:         name,
		ScreenWidth:  720,
		ScreenHeight: 1280,
		PPI:          &ppi,
		BrandID:      &brandId,
		PlatformID:   platformId}

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
	ppi := 430
	brandId := uint(2)
	platformId := uint(2)

	device := Device{
		Name:         name,
		ScreenWidth:  960,
		ScreenHeight: 1600,
		PPI:          &ppi,
		BrandID:      &brandId,
		PlatformID:   platformId}

	record := writeDb.Create(&device)

	if value, ok := record.Value.(*Device); ok {
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
