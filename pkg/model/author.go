package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
