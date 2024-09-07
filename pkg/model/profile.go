package model

import "time"

type UserProfile struct {
	User InnerContent `json:"user"`
}

type InnerContent struct {
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Following bool      `json:"following"`
	Image     string    `json:"image"`
	Token     string    `json:"token"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (up UserProfile) BuildFrom(user *User) UserProfile {
	return UserProfile{
		User: InnerContent{
			Email:     user.Email,
			Bio:       user.Bio,
			Following: user.Following,
			Image:     user.Image,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
