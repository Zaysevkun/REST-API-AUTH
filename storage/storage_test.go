package storage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres dbname=rest_api_test sslmode=disable password=4256"
	}

	os.Exit(m.Run())
}
