package sqlstorage

import (
	"database/sql"
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage"
)

// UserRepository
type UserRepository struct {
	storage *Storage
}

// Create new users table colon in db
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.storage.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id)
}

//find user in db by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.storage.db.QueryRow("SELECT  id, email, encrypted_password FROM  users WHERE  email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err == sql.ErrNoRows {
		return nil, storage.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.storage.db.QueryRow("SELECT  id, email, encrypted_password FROM  users WHERE  id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err == sql.ErrNoRows {
		return nil, storage.ErrRecordNotFound
	}

	return u, nil
}
