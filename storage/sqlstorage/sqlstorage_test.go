package sqlstorage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// get test db adress from .env
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres dbname=rest_api sslmode=disable password=4256"
	}

	os.Exit(m.Run())
}
