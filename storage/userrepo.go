package storage

import "github.com/Zaysevkun/RESTful-API/model"

type UserRepository struct {
	storage *Storage
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.storage.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.storage.db.QueryRow("SELECT  id, email, encrypted_password FROM  users WHERE  email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
