package msg

import (
	"errors"
)

type RegisterMessage struct {
	Inner InnerContent `json:"user"`
}

type InnerContent struct {
	Credentials
	Username string `json:"username"`
}

func (l *RegisterMessage) IsValid() error {
	switch {
	case l.Inner.Username == "":
		return errors.New("username is required for logon")
	default:
		return l.Inner.Credentials.IsValid()
	}
}
