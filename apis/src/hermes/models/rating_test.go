package models

import (
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateRating(t *testing.T) {
	db, dbError := database.GetWriteDB()

	if dbError != nil {
		t.Fatal("Could not get connect to write database")
	}

	defer db.Close()

	description := uniuri.New()

	rating := Rating{
		Rating:          5,
		Description:     description,
		AppVersion:      "2.0",
		PlatformVersion: "9.0",
		BrowserVersion:  "59.0",
		HasMessage:      true,
		AppID:           1,
		RangeID:         2,
		PlatformID:      3,
		AppUserID:       4,
		DeviceID:        3,
		BrowserID:       2}

	result := db.Create(&rating)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Rating); ok {
		require.Equal(t, rating.Rating, value.Rating)
		require.Equal(t, description, value.Description)
		require.Equal(t, rating.AppVersion, value.AppVersion)
		require.Equal(t, rating.PlatformVersion, value.PlatformVersion)
		require.Equal(t, rating.BrowserVersion, value.BrowserVersion)
		require.Equal(t, rating.HasMessage, value.HasMessage)
		require.Equal(t, rating.AppID, value.AppID)
		require.Equal(t, rating.RangeID, value.RangeID)
		require.Equal(t, rating.PlatformID, value.PlatformID)
		require.Equal(t, rating.AppUserID, value.AppUserID)
		require.Equal(t, rating.DeviceID, value.DeviceID)
		require.Equal(t, rating.BrowserID, value.BrowserID)
	} else {
		t.Fatal("Value is not a Rating")
	}
}
