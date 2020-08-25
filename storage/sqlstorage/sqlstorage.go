package sqlstorage

import (
	"database/sql"
	"github.com/Zaysevkun/RESTful-API/storage"
	_ "github.com/lib/pq"
)

// struct wrapper for db
type Storage struct {
	db             *sql.DB
	userRepository *UserRepository
}

// create new Storage
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// create userrepo
func (st *Storage) User() storage.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}

	st.userRepository = &UserRepository{
		storage: st,
	}

	return st.userRepository
}
