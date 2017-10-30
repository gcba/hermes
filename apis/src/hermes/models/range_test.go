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

	rangeRecord := Range{
		Name: uniuri.NewLen(11),
		From: 0,
		To:   5,
		Key:  "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())}

	result := writeDb.Create(&rangeRecord)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Range); ok {
		require.Equal(t, rangeRecord.From, value.From)
		require.Equal(t, rangeRecord.To, value.To)
		require.Equal(t, rangeRecord.Key, value.Key)
	} else {
		t.Fatal("Value is not a Range")
	}
}
