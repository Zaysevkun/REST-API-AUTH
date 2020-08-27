package sqlstorage_test

import (
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage/sqlstorage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUserRepository_Create
func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstorage.TestDb(t, databaseURL)
	defer teardown("users")
	s := sqlstorage.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// TestUserRepository_FindByEmail
func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstorage.TestDb(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstorage.TestDb(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().Find(u.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
