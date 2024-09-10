package model

import (
	"time"
)

type Article struct {
	ID              uint64
	AuthorId        uint64
	Body            string
	CreatedAt       time.Time
	Description     string
	Favorited       bool
	FavouritesCount uint32
	Slug            string
	TagList         []string
	Title           string
	UpdatedAt       time.Time
}
