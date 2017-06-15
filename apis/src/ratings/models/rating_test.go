package models

import (
	"ratings/controller"
	"testing"

	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/require"
)

func TestCreateRating(t *testing.T) {
	db := controller.GetWriteDB()
	defer db.Close()

	description := uniuri.New()

	rating := Rating{
		Rating: 5,
		Description: description,
		AppVersion: "2.0",
		PlatformVersion: "9.0",
		BrowserVersion: "59.0",
		HasMessage: true,
		AppID: 1,
		RangeID: 2,
		PlatformID: 3,
		AppUserID: 4,
		DeviceID: 5,
		BrowserID: 6
	}

	result := db.Create(&rating)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Rating); ok {
		require.Equal(t, device.Rating, value.Rating)
		require.Equal(t, content, value.Description)
		require.Equal(t, device.AppVersion, value.AppVersion)
		require.Equal(t, device.PlatformVersion, value.PlatformVersion)
		require.Equal(t, device.BrowserVersion, value.BrowserVersion)
		require.Equal(t, device.HasMessage, value.HasMessage)
		require.Equal(t, device.AppID, value.AppID)
		require.Equal(t, device.RangeID, value.RangeID)
		require.Equal(t, device.PlatformID, value.PlatformID)
		require.Equal(t, device.AppUserID, value.AppUserID)
		require.Equal(t, device.DeviceID, value.DeviceID)
		require.Equal(t, device.BrowserID, value.BrowserID)
	} else {
		t.Fatal("Value is not a Rating")
	}
}