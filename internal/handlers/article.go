package handlers

import (
	"net/http"
	"rwa/internal/params"
	"rwa/internal/service"
	"rwa/internal/utils"
	"rwa/pkg/msg"
)

type ArticleHandler struct {
	serv *service.ArticleService
}

func NewArticleHandler(userService *service.UserService) *ArticleHandler {
	return &ArticleHandler{serv: service.NewArticleService(userService)}
}

// Create an article
func (h *ArticleHandler) Create(rw http.ResponseWriter, req *http.Request) {
	article, err := utils.ReadFromRequest[*msg.CreateArticleMessage](req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	author, err := h.serv.ResolveAuthor(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	created, err := h.serv.CreateArticle(author, *article)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if respBytes, ok := utils.Marshall(rw, created); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusCreated)
	}
}

// Get recent articles globally
func (h *ArticleHandler) GetRecent(rw http.ResponseWriter, req *http.Request) {
	filter := params.GetFilterParams(req)
	result := h.serv.GetRecentGlobally(filter)
	if respBytes, ok := utils.Marshall(rw, result); ok {
		utils.SafeResponseWrite(rw, respBytes, http.StatusOK)
	}
}

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
