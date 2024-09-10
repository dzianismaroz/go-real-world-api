package repository

import (
	"bytes"
	"errors"
	"fmt"
	"rwa/internal/utils"
	"rwa/pkg/model"
	"sync"
	"time"
)

const (
	missedIdErr = "user's id required"
)

type userInMemRepository struct {
	lock sync.RWMutex
	pk   *int // primary key
	db   map[uint64]model.User
}

func NewUserRepository() *userInMemRepository {
	primarykey := 0
	return &userInMemRepository{db: make(map[uint64]model.User, 10), pk: &primarykey}
}

func (r *userInMemRepository) Add(user model.User) (model.User, error) {

	// user with such ID already exists
	if user.ID != 0 {
		return model.User{}, errors.New("already exists")
	}
	// user with email / username already exists
	for _, v := range r.db {
		if v.Equals(&user) {
			return model.User{}, errors.New("already exists")
		}
	}

	r.lock.Lock()
	*r.pk += 1
	newId := uint64(*r.pk)
	user.ID = newId
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.db[newId] = user
	r.lock.Unlock()

	return user, nil
}

func (r *userInMemRepository) Update(user model.User) (model.User, error) {
	if user.ID == 0 {
		return model.User{}, errors.New(missedIdErr)
	}
	if _, ok := r.db[user.ID]; !ok {
		return model.User{}, errors.New("No such user")
	}
	var result model.User
	r.lock.Lock()
	r.db[user.ID] = user
	result = r.db[user.ID]
	r.lock.Unlock()
	return result, nil
}

func (r *userInMemRepository) Delete(user *model.User) error {
	if user.ID == 0 {
		return errors.New(missedIdErr)
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.db, user.ID)
	return nil
}

func (r *userInMemRepository) Find(id uint64) (model.User, error) {
	noResultsErr := errors.New("no results")
	if len(r.db) < 1 {
		return model.User{}, noResultsErr
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, v := range r.db {
		if v.ID == id {
			return v, nil
		}
	}
	return model.User{}, noResultsErr
}

func (r *userInMemRepository) Authorize(logon *model.User) (model.User, error) {
	existentUser, err := r.FindBy(logon.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to logon : %w", err)
	}

	salt := string(existentUser.PasswordHash[0:8])
	if !bytes.Equal(utils.HashPass(string(logon.PasswordHash), salt), existentUser.PasswordHash) {
		return model.User{}, fmt.Errorf("Bad pass")
	}

	return existentUser, nil
}

func (r *userInMemRepository) FindBy(email string) (model.User, error) {
	noResultsErr := errors.New("no results")
	if len(r.db) < 1 {
		return model.User{}, noResultsErr
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, v := range r.db {
		if v.Email == email {
			return v, nil
		}
	}
	return model.User{}, noResultsErr
}
