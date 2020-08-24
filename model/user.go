package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// this struct corresponds with "users" table from db
type User struct {
	Id                int
	Email             string
	Password          string // not present in db
	EncryptedPassword string
}

// validate information before passing it to db(using ozzo-validation)
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(RequiredIf(u.EncryptedPassword == "")), validation.Length(8, 16)))
}

// encrypt password defore passing to db
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := EncryptMessage(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// encrypt string using bcrypt
func EncryptMessage(password string) (string, error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encryptedPass), nil
}

// helper func: if cond, validate that value cannot be blank
func RequiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
