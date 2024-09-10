package msg

import "time"

type UserProfile struct {
	Inner InnerContent `json:"user"`
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

func (up UserProfile) IsValid() bool {
	switch {
	case up.Inner.Email == "":
		return false
	default:
		return true
	}
}
