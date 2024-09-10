package repository

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	. "rwa/internal/repository/api"
	"rwa/internal/utils"
	. "rwa/pkg/model"
	"strings"
	"sync"
)

const tokenConst = "Token "

var lockSession = sync.RWMutex{}
var instance *inMemSessionManager

type inMemSessionManager struct {
	sessions       map[SessionId]Session
	sessionsOfUser map[UserId]SessionList
}

func GetSessionManager() *inMemSessionManager {
	if instance == nil {
		lockSession.Lock()
		defer lockSession.Unlock()
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

func (im *inMemSessionManager) Check(req *http.Request) (UserId, error) {
	token, err := extractTokenFrom(req)
	if err != nil {
		return 0, errors.New("no valid sessions")
	}
	lockSession.RLock()
	defer lockSession.RUnlock()
	if session, ok := im.sessions[token]; !ok {
		return 0, fmt.Errorf("no valid sessions by id %s", token)
	} else {
		return session.UserId, nil
	}
}

func extractTokenFrom(req *http.Request) (string, error) {
	token := req.Header.Get("Authorization")
	if len(token) < 1 || !strings.HasPrefix(token, tokenConst) {
		return "", errors.New("no token")
	}
	return strings.Replace(token, tokenConst, "", -1), nil
}

func (im *inMemSessionManager) Create(user User) SessionId {
	log.Println("----- creating session for user: ", user)
	session := Session{UserId: user.ID, SessionId: utils.RandStringRunes(32)}
	lockSession.Lock()
	defer lockSession.Unlock()
	im.sessions[session.SessionId] = session
	userSessions := im.sessionsOfUser[user.ID]
	userSessions = append(userSessions, session)
	im.sessionsOfUser[user.ID] = userSessions
	return session.SessionId
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
	if user.ID == 0 {
		return errors.New("invalid user")
	}
	sessions, ok := im.sessionsOfUser[user.ID]
	if !ok || len(sessions) < 1 {
		return fmt.Errorf("no active sessions by user %s", user.Username)
	}

	for _, s := range sessions {
		delete(im.sessions, s.SessionId)
	}

	delete(im.sessionsOfUser, user.ID)

	return nil
}
