package models

import (
	"ratings/controller"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/require"
)

func TestCreateAppUser(t *testing.T) {
	db := controller.GetWriteDB()
	defer db.Close()

	name := uniuri.New()
	email := "test@test.com"
	mibaID := uint(2)

	appuser := AppUser{Name: name, Email: email, MiBAID: mibaID}
	result := db.Create(&appuser)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*AppUser); ok {
		require.Equal(t, name, value.Name)
		require.Equal(t, email, value.Email)
		require.Equal(t, mibaID, value.MiBAID)
	} else {
		t.Fatal("Value is not an AppUser")
	}
}

func TestGetAppUser(t *testing.T) {
	writeDb := controller.GetWriteDB()
	defer writeDb.Close()
	readDb := controller.GetReadDB()
	defer readDb.Close()

	var result AppUser

	name := uniuri.New()
	email := "test@test.com"
	mibaID := uint(2)

	appuser := AppUser{Name: name, Email: email, MiBAID: mibaID}
	record := writeDb.Create(&appuser)

	if value, ok := record.Value.(*AppUser); ok {
		readDb.First(&result, value.ID)

		require.Equal(t, value.Name, result.Name)
		require.Equal(t, email, result.Email)
		require.Equal(t, mibaID, result.MiBAID)
	} else {
		t.Fatal("Value is not aan AppUser")
	}
}
