package storage

import "github.com/Zaysevkun/RESTful-API/model"

// using this interface we can alter witch storage implementation of storage we want to use
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
