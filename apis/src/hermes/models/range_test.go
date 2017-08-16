package models

import (
	"strings"
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetRange(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	rangeRecord := Range{
		From:  0,
		To:    5,
		Key:   "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New()),
		AppID: uint(1)}

	result := writeDb.Create(&rangeRecord)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Range); ok {
		var result Range

		readDb.First(&result, value.ID)
		require.Equal(t, rangeRecord.From, value.From)
		require.Equal(t, rangeRecord.To, value.To)
		require.Equal(t, rangeRecord.Key, value.Key)
		require.Equal(t, rangeRecord.AppID, value.AppID)
	} else {
		t.Fatal("Value is not a Range")
	}
}
