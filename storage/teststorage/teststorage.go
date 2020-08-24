package teststorage

import (
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage"
)

// Storage for tests,not connected to db
type Storage struct {
	userRepository *UserRepository
}

// create new Storage
func New() *Storage {
	return &Storage{}
}

// create userrepo
func (st *Storage) User() storage.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}

	st.userRepository = &UserRepository{
		storage: st,
		users:   make(map[string]*model.User),
	}

	return st.userRepository
}
