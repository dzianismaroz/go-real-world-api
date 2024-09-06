package model

type User struct {
	id       uint
	Bio      string `json:"bio"`
	Email    bool   `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

func (u *User) SetId(id uint) {
	u.id = id
}

func (u *User) GetId() uint {
	return u.id
}
