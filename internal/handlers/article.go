package handlers

import (
	"net/http"
	"rwa/internal/service"
)

type ArticleHandler struct {
	serv *service.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{serv: service.NewArticleService()}
}

// Create an article
func (h *ArticleHandler) Create(rw http.ResponseWriter, req *http.Request) {}

// Get recent articles globally
func (h *ArticleHandler) GetRecent(rw http.ResponseWriter, req *http.Request) {}

// Get recent articles from users you follow
func (as *ArticleHandler) GetByFollowers(rw http.ResponseWriter, req *http.Request) {}

// Update an Article
func (as *ArticleHandler) Update(rw http.ResponseWriter, req *http.Request) {}

// Get an article
func (as *ArticleHandler) GetArticle(rw http.ResponseWriter, req *http.Request) {}

// Delete an Article
func (as *ArticleHandler) DeleteArticle(rw http.ResponseWriter, req *http.Request) {}

// Get comments for an article
func (as *ArticleHandler) GetComments(rw http.ResponseWriter, req *http.Request) {}

// Create a comment for an article
func (as *ArticleHandler) PostComments(rw http.ResponseWriter, req *http.Request) {}

// Delete comments for an article
func (as *ArticleHandler) DeleteComments(rw http.ResponseWriter, req *http.Request) {}
