package repository

import "rwa/pkg/model"

type ArticilesRepository interface {
	GetAllArticles() []model.Article
}
