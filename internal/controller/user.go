package service

import (
	"rwa/internal/repository"
	"rwa/pkg/model"
	models "rwa/pkg/model"
)

type IUserService interface {
	Register(user models.User) (models.User, error)
	GetCurrent() (models.User, error)
	Login() error
	Logout() error
}

type userDBService struct {
	repository *repository.UserRepository
}

func NewUserController(repository *repository.UserRepository) *userDBService {
	return &userDBService{repository: repository}
}

func (uc *userDBService) Register(user models.User) (model.User, error) {
	return models.User{}, nil
}

func (uc *userDBService) GetCurrent(user models.User) (model.User, error) {
	return models.User{}, nil
}

func (uc *userDBService) Login() error {
	return nil
}

func (uc *userDBService) Logout() error {
	return nil
}
