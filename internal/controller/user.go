package service

import (
	"fmt"
	"net/http"
	api "rwa/internal/repository/api"
	impl "rwa/internal/repository/inmemory"
	"rwa/internal/utils"
	"rwa/pkg/model"
	m "rwa/pkg/model/msg"
	"time"
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
	if err != nil {
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

func (uc *userInMemService) GetCurrent(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("current use is unknown yet")); err != nil {
		http.Error(rw, "failed to register", http.StatusInternalServerError)
	}
}

func (uc *userInMemService) UpdateCurrent(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("current use is unknown yet")); err != nil {
		http.Error(rw, "failed to register", http.StatusInternalServerError)
	}
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

	sessID := utils.RandStringRunes(32)
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sessID,
		Expires: time.Now().Add(90 * 24 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(rw, cookie)
	respBytes, ok := utils.Marshall(rw, profile)
	if !ok {
		return
	}
	utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
}

func (uc *userInMemService) Logout() error {
	return nil
}
