package api

import (
	"rwa/internal/params"
	"rwa/pkg/model"
)

type ArticilesRepository interface {
	GetAllArticles(filter params.FilterParams) []model.Article
	Save(article model.Article) (model.Article, error)
}
