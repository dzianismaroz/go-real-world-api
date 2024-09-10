package handlers

import (
	"fmt"
	"net/http"
	"rwa/internal/converter"
	repository "rwa/internal/repository/inmemory"
	s "rwa/internal/service"
	"rwa/internal/utils"
	"rwa/pkg/msg"
)

type UserHandler struct {
	service *s.UserService
}

func NewUserHandler(service *s.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register a new user
func (h *UserHandler) Register(rw http.ResponseWriter, req *http.Request) {
	registerMessage, err := utils.ReadFromRequest[msg.RegisterMessage](req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// attempt to register new user
	createdUser, err := h.service.Register(registerMessage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// reply with newly created user's profile

	profile := converter.ToProfile(createdUser)
	token := repository.GetSessionManager().Create(createdUser)
	profile.Inner.Token = token
	if respBytes, ok := utils.Marshall(rw, profile); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusCreated)
	}
}

// Get current user
func (h *UserHandler) GetCurrent(rw http.ResponseWriter, req *http.Request) {
	existentUser, err := h.service.GetCurrent(req)
	if err != nil {
		http.Error(rw, "failed to resolve curent user", http.StatusUnprocessableEntity)
		return
	}
	profile := converter.ToProfile(existentUser)
	if respBytes, ok := utils.Marshall(rw, profile); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
	}
}

// Update current user
//
//nolint:errcheck
func (h *UserHandler) UpdateCurrent(rw http.ResponseWriter, req *http.Request) {
	updateMessage, err := utils.ReadFromRequest[msg.UserProfile](req)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to update user : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}
	result, err := h.service.Update(req, updateMessage)
	if err != nil {

		http.Error(rw, fmt.Errorf("failed to update user : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}

	profile := converter.ToProfile(result)
	sm := repository.GetSessionManager()
	sm.DestroyCurrent(rw, req)
	token := sm.Create(result)
	profile.Inner.Token = token
	if respBytes, ok := utils.Marshall(rw, profile); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
	}
}

// Existing user login
func (h *UserHandler) Login(rw http.ResponseWriter, req *http.Request) {
	logonMessage, err := utils.ReadFromRequest[msg.LogonMessage](req)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to logon : %w", err).Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := h.service.Login(logonMessage)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to logon : %w", err).Error(), http.StatusUnauthorized)
		return
	}

	profile := converter.ToProfile(user)
	token := repository.GetSessionManager().Create(user)
	profile.Inner.Token = token

	if respBytes, ok := utils.Marshall(rw, profile); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
	}
}

// Existing user logout
func (h *UserHandler) Logout(rw http.ResponseWriter, req *http.Request) {
	if err := repository.GetSessionManager().DestroyCurrent(rw, req); err != nil {
		http.Error(rw, fmt.Errorf("failed to logout : %w", err).Error(), http.StatusInternalServerError)
		return
	}
	utils.SafeResponseWrite(rw, nil, http.StatusOK)
}
