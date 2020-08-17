package model_test

import (
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	err := u.BeforeCreate()

	assert.NoError(t, err)
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				usr := model.TestUser(t)
				usr.Email = ""
				return usr
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				usr := model.TestUser(t)
				usr.Email = "aye"
				return usr
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				usr := model.TestUser(t)
				usr.Password = ""
				return usr
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				usr := model.TestUser(t)
				usr.Password = "aye"
				return usr
			},
			isValid: false,
		},
		{
			name: "empty password but not empty encPassword",
			u: func() *model.User {
				usr := model.TestUser(t)
				usr.Password = ""
				usr.EncryptedPassword = "encryptedPassword"
				return usr
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}

}
