package utils

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetConfig(variable string, fallback string) string {
	if result := os.Getenv(variable); len(result) > 0 {
		return result
	}

	return fallback
}

func GetPort(variable string, defaultPort int) int {
	port, portErr := strconv.Atoi(os.Getenv(variable))

	if portErr != nil || port == 0 {
		return defaultPort
	}

	return port
}

func GetPublicKey(variable string, echo *echo.Echo) *rsa.PublicKey {
	rawKey, readErr := ioutil.ReadFile(os.Getenv(variable))

	if readErr != nil {
		echo.Logger.Fatal("Could not find key file")
	}

	key, parseErr := jwt.ParseRSAPublicKeyFromPEM(rawKey)

	if parseErr != nil {
		echo.Logger.Fatal(parseErr.Error())
	}

	return key
}
