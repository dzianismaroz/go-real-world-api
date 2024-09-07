package msg

import "errors"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) IsValid() error {
	switch {
	case c.Email == "":
		return errors.New("email is required for logon")
	case c.Password == "":
		return errors.New("password is required for logon")
	default:
		return nil
	}
}
