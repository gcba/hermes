package models

import (
	"testing"

	"hermes/database"

	"github.com/dchest/uniuri"
	_ "github.com/joho/godotenv/autoload" // Loads config from .env file
	"github.com/stretchr/testify/require"
)

func TestCreateMessage(t *testing.T) {
	db, dbError := database.GetWriteDB()

	if dbError != nil {
		t.Fatal("Could not get connect to write database")
	}

	defer db.Close()

	content := uniuri.New()

	message := Message{
		Message:   content,
		Direction: "in",
		RatingID:  1}

	result := db.Create(&message)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Message); ok {
		require.Equal(t, content, value.Message)
		require.Equal(t, message.Direction, value.Direction)
		require.Equal(t, message.RatingID, value.RatingID)
	} else {
		t.Fatal("Value is not a Message")
	}
}
