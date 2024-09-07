package repository

import (
	"errors"
	"fmt"
	"net/http"
	. "rwa/internal/repository/api"
	. "rwa/pkg/model"
	"sync"
)

var lock = sync.Mutex{}
var instance *inMemSessionManager

type inMemSessionManager struct {
	sessions       map[SessionId]Session
	sessionsOfUser map[UserId]SessionList
}

func GetSessionManager() *inMemSessionManager {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = newInMemSessionManager()
	}
	return instance
}

func newInMemSessionManager() *inMemSessionManager {
	return &inMemSessionManager{
		sessions:       make(map[SessionId]Session, 10),
		sessionsOfUser: make(map[UserId]SessionList, 10),
	}
}

func (im *inMemSessionManager) Check(req *http.Request) error {
	token, err := extractTokenFrom(req)
	if err != nil {
		return errors.New("no valid sessions")
	}
	if _, ok := im.sessions[token]; !ok {
		return fmt.Errorf("no valid sessions by id %s", token)
	} else {
		return nil
	}
}

func extractTokenFrom(req *http.Request) (string, error) {
	token := req.Header.Get("Authorization")
	if len(token) < 1 {
		return "", errors.New("no token")
	}
	return token, nil
}

func (im *inMemSessionManager) Create(rw http.ResponseWriter, user *User) error {
	panic("not implemented") // TODO: Implement
}

func (im *inMemSessionManager) DestroyCurrent(rw http.ResponseWriter, req *http.Request) error {
	sessionId, err := extractTokenFrom(req)
	if err != nil {
		return errors.New("no valid sessions")
	}
	resolvedSession, ok := im.sessions[sessionId]
	if !ok {
		return errors.New("no valid session")
	}
	delete(im.sessions, sessionId)
	sessions, ok := im.sessionsOfUser[resolvedSession.UserId]
	if ok {
		cleanSessions := make(SessionList, 0, len(sessions))
		for _, s := range sessions {
			if s.SessionId != sessionId {
				cleanSessions = append(cleanSessions, s)
			}
		}
		im.sessionsOfUser[resolvedSession.UserId] = cleanSessions
	}
	return nil
}

func (im *inMemSessionManager) DestroyAll(rw http.ResponseWriter, user *User) error {
	if user.GetId() == 0 {
		return errors.New("invalid user")
	}
	sessions, ok := im.sessionsOfUser[user.GetId()]
	if !ok || len(sessions) < 1 {
		return fmt.Errorf("no active sessions by user %s", user.Username)
	}

	for _, s := range sessions {
		delete(im.sessions, s.SessionId)
	}

	delete(im.sessionsOfUser, user.GetId())

	return nil
}
