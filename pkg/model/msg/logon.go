package msg

// LogonMessage is login request.
type LogonMessage struct {
	Inner Credentials `json:"user"`
}

func (l *LogonMessage) IsValid() error {
	return l.Inner.IsValid()
}
