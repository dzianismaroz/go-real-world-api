package service

import (
	"net/http"
	api "rwa/internal/repository/api"
	. "rwa/internal/repository/inmemory"
)

type ArticleService struct {
	repo api.ArticilesRepository
}

func NewArticleService() *ArticleService {
	return &ArticleService{repo: NewArticlesRepository()}
}

func (as *ArticleService) CreateArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusCreated)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) GetRecentGlobally(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) GetRecentFollowers(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) UpdateArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) GetArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) DeleteArticle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) GetComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) PostComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}

func (as *ArticleService) DeleteComments(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ARTICLES"))
	if err != nil {
		http.Error(rw, "ERROR", http.StatusInternalServerError)
	}
}
