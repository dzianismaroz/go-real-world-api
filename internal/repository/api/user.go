package api

import (
	. "rwa/pkg/model"
)

type UserRepository interface {
	Authorize(logon *User) (User, error)
	Find(id uint64) (User, error)
	FindBy(email string) (User, error)
	Add(user User) (User, error)
	Update(user User) (User, error)
	Delete(user *User) error
}
