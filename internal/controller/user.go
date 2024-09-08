package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	auth "rwa/internal/handlers/http"
	api "rwa/internal/repository/api"
	impl "rwa/internal/repository/inmemory"
	repository "rwa/internal/repository/inmemory"
	"rwa/internal/utils"
	"rwa/pkg/model"
	m "rwa/pkg/model/msg"
)

type userInMemService struct {
	repository api.UserRepository
}

func NewUserController() *userInMemService {
	return &userInMemService{repository: impl.NewUserRepository()}
}

// Register new user from RegisterMessage. Reply a UserProfile
func (uc *userInMemService) Register(rw http.ResponseWriter, req *http.Request) {
	// extract request payload to RegisterMessage
	registerMessage, err := utils.ReadFromRequest[m.RegisterMessage](req)
	if err != nil || !registerMessage.IsValid() {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	// attempt to register new user
	createdUser, err := uc.repository.Add(registerMessage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// reply with newly created user's profile
	profile := model.UserProfile{}.BuildFrom(createdUser)
	respBytes, ok := utils.Marshall(rw, profile)
	if !ok {
		return
	}
	utils.SafeResponseWrite(rw, respBytes, http.StatusCreated)
}

func (uc *userInMemService) resolveCurrent(req *http.Request) (*model.User, error) {
	userIdRaw := req.Context().Value(auth.UserKey)
	if userIdRaw == 0 {
		return nil, errors.New("failed to resolve curent user")
	}
	userID := userIdRaw.(uint64)
	existentUser, err := uc.repository.Find(userID)
	if err != nil {
		return nil, errors.New("failed to resolve curent user")
	}
	return existentUser, nil
}

func (uc *userInMemService) GetCurrent(rw http.ResponseWriter, req *http.Request) {
	existentUser, err := uc.resolveCurrent(req)
	if err != nil {
		http.Error(rw, "failed to resolve curent user", http.StatusUnprocessableEntity)
		return
	}
	profile := model.UserProfile{}.BuildFrom(existentUser)
	respBytes, ok := utils.Marshall(rw, profile)
	if !ok {
		return
	}
	utils.SafeResponseWrite(rw, respBytes, http.StatusOK)

}

func (uc *userInMemService) UpdateCurrent(rw http.ResponseWriter, req *http.Request) {
	existentUser, err := uc.resolveCurrent(req)
	if err != nil {
		http.Error(rw, "failed to resolve curent user", http.StatusUnprocessableEntity)
		return
	}
	updateMessage, err := utils.ReadFromRequest[model.UserProfile](req)
	if err != nil {
		log.Println("#### is valid ?", updateMessage.IsValid())
		http.Error(rw, fmt.Errorf("failed to update user : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}
	result, err := uc.repository.Update(existentUser, updateMessage)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to update user : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}

	profile := model.UserProfile{}.BuildFrom(result)
	sm := repository.GetSessionManager()
	sm.DestroyCurrent(rw, req)
	token := sm.Create(result)
	profile.Inner.Token = token
	respBytes, ok := utils.Marshall(rw, profile)
	if !ok {
		return
	}
	utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
}

func (uc *userInMemService) Login(rw http.ResponseWriter, req *http.Request) {
	logonMessage, err := utils.ReadFromRequest[m.LogonMessage](req)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to logon : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}

	profile, err := uc.repository.Authorize(logonMessage)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to logon : %w", err).Error(), http.StatusUnauthorized)
		return
	}

	respBytes, ok := utils.Marshall(rw, profile)
	if !ok {
		return
	}
	utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
}

func (uc *userInMemService) Logout() error {
	return nil
}
