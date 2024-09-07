package model

type Author struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
