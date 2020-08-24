package teststorage

import (
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage"
)

//UserRepository
type UserRepository struct {
	storage *Storage
	users   map[string]*model.User
}

// create new user in fake db
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.Id = len(r.users)

	return nil
}

// find user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}

	return u, nil
}
