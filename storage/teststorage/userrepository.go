package teststorage

import (
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage"
)

//UserRepository
type UserRepository struct {
	storage *Storage
	users   map[int]*model.User
}

// create new user in fake db
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

// find user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, storage.ErrRecordNotFound
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}

	return u, nil
}
