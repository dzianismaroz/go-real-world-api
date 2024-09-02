package model

type User struct {
	Id       int    `json:"-"`
	Bio      string `json:"bio"`
	Email    bool   `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
