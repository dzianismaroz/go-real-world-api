package msg

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Credentials) IsValid() bool {
	switch {
	case c.Email == "":
		return false
	case c.Password == "":
		return false
	default:
		return true
	}
}
