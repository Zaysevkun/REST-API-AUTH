package sqlstorage

import (
	"database/sql"
	"log"
	"strings"
	"testing"
)

// test example of db with teardown func
func TestDb(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			arg := strings.Join(tables, ", ")
			_, err := db.Exec("TRUNCATE " + arg + " CASCADE")
			log.Println(err)
		}

		db.Close()
	}
}
