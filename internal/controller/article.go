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

func (as *articleService) CreateArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusCreated)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) GetRecentGlobally(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) GetRecentFollowers(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) UpdateArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) GetArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) DeleteArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) GetComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) PostComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *articleService) DeleteComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}
