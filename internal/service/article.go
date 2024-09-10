package service

import (
	"log"
	"net/http"
	"rwa/internal/converter"
	"rwa/internal/params"
	api "rwa/internal/repository/api"
	repository "rwa/internal/repository/inmemory"
	"rwa/pkg/model"
	"rwa/pkg/msg"
	"slices"
)

type ArticleService struct {
	repo        api.ArticilesRepository
	userService *UserService
}

func NewArticleService(userService *UserService) *ArticleService {
	return &ArticleService{
		repo:        repository.NewArticlesRepository(),
		userService: userService}
}

func (as *ArticleService) ResolveAuthor(req *http.Request) (model.User, error) {
	return as.userService.ResolveCurrent(req)
}

func (as *ArticleService) CreateArticle(author model.User, articleRequest *msg.CreateArticleMessage) (msg.Article, error) {
	article := converter.ToArticleEntity(articleRequest, author.ID)
	saved, err := as.repo.Save(article)
	if err != nil {
		log.Println("failed to create article:", err)
		return msg.Article{}, err
	}

	result := converter.ToArticleMsg(saved, converter.ToProfile(author))
	return result, nil
}

func (as *ArticleService) GetRecentGlobally(filter params.FilterParams) msg.ArticlesList {
	articles := as.repo.GetAllArticles(filter)
	result := make([]msg.ArticleBody, 0, len(articles))
	for _, a := range articles {
		author := as.userService.Find(a.AuthorId)
		if filter.HasAuthorFilter() {
			if author.Username != filter.Author {
				continue
			}
		}
		if filter.HasTagFilter() {
			if !slices.Contains(a.TagList, filter.Tag) {
				continue
			}
		}
		result = append(result, converter.ToArticleBodyMsg(a, converter.ToProfile(author)))
	}
	return converter.ToArticlesList(result)
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
