package repository

import "rwa/pkg/model"

type UserRepository interface {
	Save(user model.User)   //(model.User, error)
	Update(user model.User) //(model.User, error)
	Delete(user model.User) //error
}

type UserDBRepository struct{}

func NewUserRepository() *UserRepository {
	return nil
}
