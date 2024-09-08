package model

import (
	"rwa/pkg/model/msg"
	"time"
)

type User struct {
	id           uint64
	Email        string
	Username     string
	Bio          string
	Image        string
	Following    bool
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) SetId(id uint64) {
	u.id = id
}

func (u *User) GetId() uint64 {
	return u.id
}

func (u *User) IsSameIdentity(other *User) bool {
	return u.Email == other.Email || u.Username == other.Username
}

func (u User) BuildFrom(r *msg.RegisterMessage) *User {
	return &User{
		Username: r.Inner.Username,
		Email:    r.Inner.Email,
	}
}

func (u *User) MergeFrom(other *UserProfile) User {
	result := User{
		id:           u.id,
		Email:        u.Email,
		Bio:          u.Bio,
		Username:     u.Username,
		Image:        u.Image,
		Following:    u.Following,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	if other.Inner.Username != "" {
		result.Username = other.Inner.Username
	}

	if other.Inner.Email != "" {
		result.Email = other.Inner.Email
	}
	result.Bio = other.Inner.Bio
	return result
}
