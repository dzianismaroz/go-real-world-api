package service

import (
	"errors"
	"net/http"
	"rwa/internal/converter"
	api "rwa/internal/repository/api"
	repository "rwa/internal/repository/inmemory"
	"rwa/internal/utils"
	"rwa/pkg/model"
	"rwa/pkg/msg"
)

type UserService struct {
	repository api.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repository: repository.NewUserRepository()}
}

// Register new User
func (s *UserService) Register(candidate *msg.RegisterMessage) (model.User, error) {
	return s.repository.Add(converter.ToUser(candidate))
}

func (s *UserService) Find(userId uint64) model.User {
	result, err := s.repository.Find(userId)
	if err != nil {
		return model.User{}
	}
	return result
}

func (s *UserService) ResolveCurrent(req *http.Request) (model.User, error) {
	userIdRaw := req.Context().Value(utils.UserKey)
	if userIdRaw == 0 {
		return model.User{}, errors.New("failed to resolve curent user")
	}
	userID := userIdRaw.(uint64)
	existentUser, err := s.repository.Find(userID)
	if err != nil {
		return model.User{}, errors.New("failed to resolve curent user")
	}
	return existentUser, nil
}

func (s *UserService) GetCurrent(req *http.Request) (model.User, error) {
	return s.ResolveCurrent(req)
}

func (s *UserService) Update(req *http.Request, updateMessage *msg.UserProfile) (model.User, error) {
	curentUser, err := s.ResolveCurrent(req)
	if err != nil {
		return model.User{}, err
	}
	candidate := converter.Merge(curentUser, updateMessage)
	return s.repository.Update(candidate)
}

func (s *UserService) Login(logonMessage *msg.LogonMessage) (model.User, error) {
	targetUser := converter.FromLogon(logonMessage)
	return s.repository.Authorize(targetUser)
}
