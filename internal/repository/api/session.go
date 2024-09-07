package repository

import (
	"net/http"
	. "rwa/pkg/model"
)

type (
	SessionId   = string
	UserId      = uint64
	SessionList = []Session

	Session struct {
		UserId    UserId
		SessionId SessionId
	}

	SessionRepository interface {
		Check(*http.Request) (*Session, error)
		Create(http.ResponseWriter, *User) error
		DestroyCurrent(http.ResponseWriter, *http.Request) error
		DestroyAll(http.ResponseWriter, *User) error
	}
)
