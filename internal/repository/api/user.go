package repository

import . "rwa/pkg/model"

type UserRepository interface {
	Find(id uint) (*User, error)
	Add(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
}
