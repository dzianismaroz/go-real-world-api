package model

import (
	"time"
)

type User struct {
	ID           uint64
	Email        string
	Username     string
	Bio          string
	Image        string
	Following    bool
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u User) Equals(other *User) bool {
	return u.Email == other.Email || u.Username == other.Username
}
