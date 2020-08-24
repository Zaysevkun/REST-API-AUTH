package api

import (
	"database/sql"
	"github.com/Zaysevkun/RESTful-API/storage/sqlstorage"
	"net/http"
)

// start server using inputed config
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	storage := sqlstorage.New(db)
	srv := NewServer(storage)

	return http.ListenAndServe(config.Port, srv)
}

// make and check db connection on inputed db URL
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
