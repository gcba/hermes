package models

import (
	"ratings/controller"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBrand(t *testing.T) {
	db := controller.GetWriteDB()
	defer db.Close()

	brand := Brand{Name: "Test"}
	result := db.Create(&brand)

	require.Equal(t, nil, result.Error)

	if value, ok := result.Value.(*Brand); ok {
		require.Equal(t, "Test", value.Name)
	} else {
		t.Fatal("Value is not a Brand")
	}
}
