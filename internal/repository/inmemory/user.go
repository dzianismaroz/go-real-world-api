package repository

import (
	"bytes"
	"errors"
	"fmt"
	"rwa/internal/utils"
	"rwa/pkg/model"
	"time"
)

const (
	missedIdErr = "user's id required"
)

type userInMemRepository struct {
	pk uint64 // primary key
	db map[uint64]*model.User
}

func NewUserRepository() *userInMemRepository {
	return &userInMemRepository{db: make(map[uint64]*model.User, 10)}
}

func (r userInMemRepository) Add(user *model.User) (*model.User, error) {
	// user with such ID already exists
	if user.ID != 0 {
		return nil, errors.New("already exists")
	}
	// user with email / username already exists
	for _, v := range r.db {
		if v.IsSameIdentity(user) {
			return nil, errors.New("already exists")
		}
	}
	r.pk++
	r.db[r.pk] = user
	user.ID = r.pk
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return user, nil
}

func (r userInMemRepository) Update(user *model.User) (*model.User, error) {
	if user.ID == 0 {
		return nil, errors.New(missedIdErr)
	}
	if _, ok := r.db[user.ID]; !ok {
		return nil, errors.New("No such user")
	}
	r.db[user.ID] = user
	return user, nil
}

func (r userInMemRepository) Delete(user *model.User) error {
	if user.ID == 0 {
		return errors.New(missedIdErr)
	}
	delete(r.db, user.ID)
	return nil
}

func (r userInMemRepository) Find(id uint64) (*model.User, error) {
	noResultsErr := errors.New("no results")
	if len(r.db) < 1 {
		return nil, noResultsErr
	}
	for _, v := range r.db {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, noResultsErr
}

func (r userInMemRepository) Authorize(logon *model.User) (*model.User, error) {
	// log.Println("auth as : ", logon)
	existentUser, err := r.FindBy(logon.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to logon : %w", err)
	}

	salt := string(existentUser.PasswordHash[0:8])
	if !bytes.Equal(utils.HashPass(string(logon.PasswordHash), salt), existentUser.PasswordHash) {
		return nil, fmt.Errorf("Bad pass")
	}

	return existentUser, nil
}

func (r userInMemRepository) FindBy(email string) (*model.User, error) {
	noResultsErr := errors.New("no results")
	if len(r.db) < 1 {
		return nil, noResultsErr
	}
	for _, v := range r.db {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, noResultsErr
}
