package db

import "github.com/Zaysevkun/RESTful-API/model"

type UserRepository struct {
	storage Storage
}

func (r *UserRepository) Create(m *model.User) (*model.User, error) {
	return nil, nil
}
