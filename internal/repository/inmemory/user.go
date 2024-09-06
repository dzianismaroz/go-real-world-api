package repository

import (
	"errors"
	. "rwa/pkg/model"
)

const (
	missedIdErr = "user's id required"
)

type userInMemRepository struct {
	pk uint
	db map[uint]*User
}

func NewUserRepository() *userInMemRepository {
	return &userInMemRepository{db: make(map[uint]*User, 10)}
}

func (r userInMemRepository) Add(user *User) (*User, error) {
	// user with such ID already exists
	if user.GetId() != 0 {
		return nil, errors.New("user with such id already exists")
	}
	// user with email / username already exists
	for _, v := range r.db {
		if v.Email == user.Email || v.Username == user.Username {
			return nil, errors.New("already exists")
		}
	}
	r.pk++
	r.db[r.pk] = user
	user.SetId(r.pk)
	return nil, nil
}

func (r userInMemRepository) Update(user *User) (*User, error) {

	if user.GetId() == 0 {
		return nil, errors.New(missedIdErr)
	}
	r.db[user.GetId()] = user
	return user, nil
}

func (r userInMemRepository) Delete(user *User) error {
	if user.GetId() == 0 {
		return errors.New(missedIdErr)
	}
	delete(r.db, user.GetId())
	return nil
}

func (r userInMemRepository) Find(id uint) (*User, error) {
	noResultsErr := errors.New("no results")
	if len(r.db) < 1 {
		return nil, noResultsErr
	}
	for _, v := range r.db {
		if v.GetId() == id {
			return v, nil
		}
	}
	return nil, noResultsErr
}
