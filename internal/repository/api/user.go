package repository

import (
	. "rwa/pkg/model"
	model "rwa/pkg/model/msg"
)

type UserRepository interface {
	Authorize(logon *model.LogonMessage) (UserProfile, error)
	Find(id uint64) (*User, error)
	FindBy(email string) (*User, error)
	Add(user *model.RegisterMessage) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
}
