package model

type UserProfile struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}

func BuildFrom(user *User) UserProfile {
	return UserProfile{}
}
