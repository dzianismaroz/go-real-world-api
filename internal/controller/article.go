package service

import (
	"net/http"
	api "rwa/internal/repository/api"
	. "rwa/internal/repository/inmemory"
)

type articleService struct {
	repo api.ArticilesRepository
}

func NewArticleService() *articleService {
	return &articleService{repo: NewArticlesRepository()}
}

func (as *articleService) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}
