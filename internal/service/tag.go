package service

import (
	. "rwa/internal/repository/inmemory"
)

type tagsController struct {
	repo TagRepository
}

func NewTagsController() *tagsController {
	return &tagsController{repo: NewTagRepository()}
}

// List all available tags
func (tc *tagsController) ListTags() []string {
	return []string{}
}
