package repository

import (
	"rwa/pkg/model"
)

type articleRepository struct {
}

func (a *articleRepository) GetAllArticles() []model.Article {
	return nil
}

func NewArticlesRepository() *articleRepository {
	return &articleRepository{}
}
