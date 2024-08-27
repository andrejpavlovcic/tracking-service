package db

import (
	"fmt"
	"os"
	"testing"
)

func TestGetDatabaseStringConnection(t *testing.T) {
	var (
		dbHost                   = "test_host"
		dbPort                   = "5432"
		dbUser                   = "test"
		dbPassword               = "test"
		dbName                   = "test"
		expectedConnectionString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
		)
	)

	os.Setenv("DB_HOST", dbHost)
	os.Setenv("DB_PORT", dbPort)
	os.Setenv("DB_USER", dbUser)
	os.Setenv("DB_PASSWORD", dbPassword)
	os.Setenv("DB_NAME", dbName)

	gotConnectionString := getDatabaseStringConnection()

	if gotConnectionString != expectedConnectionString {
		t.Fail()
	}
}
