package msg

import "time"

type Article struct {
	Content ArticleBody `json:"article"`
}

type ArticleBody struct {
	Author          InnerContentShort `json:"author"`
	Body            string            `json:"body"`
	CreatedAt       time.Time         `json:"createdAt"`
	Description     string            `json:"description"`
	Favourited      bool              `json:"favourited"`
	FavouritedCount uint32            `json:"favouritedCount"`
	Slug            string            `json:"slug"`
	Tags            []string          `json:"tagList"`
	Title           string            `json:"title"`
	UpdatedAt       time.Time         `json:"updatedAt"`
}

type InnerContentShort struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}

func (a *Article) IsValid() bool {
	return true
}
