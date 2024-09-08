package msg

// LogonMessage is login request.
type LogonMessage struct {
	Inner Credentials `json:"user"`
}

func (l *LogonMessage) IsValid() bool {
	return l.Inner.IsValid()
}
