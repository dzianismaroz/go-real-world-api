package msg

type RegisterMessage struct {
	Inner InnerContent `json:"user"`
}

type InnerContent struct {
	Credentials
	Username string `json:"username"`
}

func (l *RegisterMessage) IsValid() bool {
	switch {
	case l.Inner.Username == "":
		return false
	default:
		return l.Inner.Credentials.IsValid()
	}
}
