package models

import (
	"ratings/database"
	"strings"
	"testing"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestGetRange(t *testing.T) {
	writeDb := database.GetWriteDB()
	defer writeDb.Close()
	readDb := database.GetReadDB()
	defer readDb.Close()

	from := 0
	to := 5
	key := "7C6F0035B18C3D5J" + strings.ToUpper(uniuri.New())

	rangeRecord := Range{From: 0, To: 5, Key: key}
	result := writeDb.Create(&rangeRecord)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Range); ok {
		var result Range

		readDb.First(&result, value.ID)
		require.Equal(t, from, value.From)
		require.Equal(t, to, value.To)
		require.Equal(t, key, value.Key)
	} else {
		t.Fatal("Value is not a Range")
	}
}
