package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	api "rwa/internal/repository/api"
	impl "rwa/internal/repository/inmemory"
	"rwa/pkg/model"

	"strings"
)

type userInMemService struct {
	repository api.UserRepository
}

func NewUserController() *userInMemService {
	return &userInMemService{repository: impl.NewUserRepository()}
}

func (uc *userInMemService) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if strings.HasSuffix(req.URL.Path, "/login") {
			uc.Login(rw, req)
			return
		}
		uc.Register(rw, req)
	case http.MethodGet:
		uc.GetCurrent(rw, req)
	}
}

func closeResources(closer io.Closer) {

	err := closer.Close()
	if err != nil {
		log.Fatal("unable to close request body")
	}
}

func safeWrite(w http.ResponseWriter, statusCode int, content []byte) {
	w.WriteHeader(statusCode)
	_, err := w.Write(content)
	if err != nil {
		http.Error(w, "unexpected error", http.StatusInternalServerError)
	}
}

func (uc *userInMemService) Register(rw http.ResponseWriter, req *http.Request) {
	rawBodyBytes, err := io.ReadAll(req.Body)
	defer closeResources(req.Body)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to register : %w", err).Error(), http.StatusInternalServerError)
		return
	}
	result := &model.User{}
	if err := json.Unmarshal(rawBodyBytes, &result); err != nil {
		http.Error(rw, fmt.Errorf("failed to register : %w", err).Error(), http.StatusInternalServerError)
		return
	}

	if _, err := uc.repository.Add(result); err != nil {
		http.Error(rw, "failed to register", http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to register : %w", err).Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	n, err := rw.Write(respBytes)
	if err != nil || n < 1 {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
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
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("attempt to login")); err != nil {
		http.Error(rw, "failed to login", http.StatusInternalServerError)
	}
}

func (uc *userInMemService) Logout() error {
	return nil
}
