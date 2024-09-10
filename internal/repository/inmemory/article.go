package repository

import (
	"rwa/internal/params"
	"rwa/pkg/model"
	"sync"
)

type slug = string
type articleRepository struct {
	mu sync.RWMutex
	db map[slug][]model.Article
}

func (a *articleRepository) GetAllArticles(filter params.FilterParams) []model.Article {
	result := make([]model.Article, 0, 10)
	for _, art := range a.db {
		result = append(result, art...)
	}
	return result
}

func NewArticlesRepository() *articleRepository {
	return &articleRepository{
		db: make(map[string][]model.Article, 10),
	}
}

func (a *articleRepository) Save(article model.Article) (model.Article, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	articles := a.db[article.Slug]
	articles = append(articles, article)
	a.db[article.Slug] = articles
	return article, nil
}
