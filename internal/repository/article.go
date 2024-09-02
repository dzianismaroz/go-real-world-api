package repository

import "rwa/pkg/model"

type ArticilesRepository interface {
	GetAllArticles() []model.Article
}

type ArticlesRepository struct {
}

func NewArticlesRepository() *ArticlesRepository {
	return nil
}
