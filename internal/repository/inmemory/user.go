package repository

import (
	"bytes"
	"errors"
	"fmt"
	"rwa/internal/utils"
	. "rwa/pkg/model"
	model "rwa/pkg/model/msg"
	"time"

	"golang.org/x/crypto/argon2"
)

const (
	missedIdErr = "user's id required"
)

type userInMemRepository struct {
	pk uint64 // primary key
	db map[uint64]*User
}

func NewUserRepository() *userInMemRepository {
	return &userInMemRepository{db: make(map[uint64]*User, 10)}
}

func (r userInMemRepository) Add(register *model.RegisterMessage) (*User, error) {
	user := User{}.BuildFrom(register)
	// user with such ID already exists
	if user.GetId() != 0 {
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
	user.SetId(r.pk)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.PasswordHash = hashPass(register.Inner.Password, utils.RandStringRunes(8))
	return user, nil
}

func hashPass(plainPassword, salt string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)
	res := make([]byte, len(salt))
	copy(res, salt)
	return append(res, hashedPass...)
}

func (r userInMemRepository) Update(user *User) (*User, error) {

	if user.GetId() == 0 {
		return nil, errors.New(missedIdErr)
	}
	if _, ok := r.db[user.GetId()]; !ok {
		return nil, errors.New("No such user")
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

func (r userInMemRepository) Find(id uint64) (*User, error) {
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

func (r userInMemRepository) Authorize(logon *model.LogonMessage) (UserProfile, error) {
	profile := UserProfile{}
	existentUser, err := r.FindBy(logon.Inner.Email)
	if err != nil {
		return profile, fmt.Errorf("failed to logon : %w", err)
	}

	salt := string(existentUser.PasswordHash[0:8])
	if !bytes.Equal(hashPass(logon.Inner.Password, salt), existentUser.PasswordHash) {
		return profile, fmt.Errorf("Bad pass")
	}
	GetSessionManager().Create()
	return profile.BuildFrom(existentUser), nil
}

func (r userInMemRepository) FindBy(email string) (*User, error) {
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
