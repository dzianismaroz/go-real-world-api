package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Author          Author    `json:"author"`
	CreatedAt       time.Time `json:"createdAt"`
	Description     string    `json:"descrtiption"`
	Favorited       bool      `json:"favourited"`
	FavouritesCount int       `json:"favouritesCount"`
	Slug            string    `json:"slug"`
	TagList         []Tag     `json:"tagList"`
	Title           string    `json:"title"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
