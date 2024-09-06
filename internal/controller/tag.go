package service

import (
	"fmt"
	"net/http"
	. "rwa/internal/repository/inmemory"
)

type tagsController struct {
	repo TagRepository
}

func NewTagsController() *tagsController {
	return &tagsController{repo: NewTagRepository()}
}

func (tc *tagsController) ListTags(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	//nolint:errcheck
	rw.Write([]byte("listing tags"))
}

func (tc *tagsController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("LIST ALL TAGS")); err != nil {
		http.Error(rw, fmt.Errorf("failed to list all tags %w", err).Error(), http.StatusInternalServerError)
	}
}
