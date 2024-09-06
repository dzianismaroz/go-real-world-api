package repository

type TagRepository interface {
}

type tagRepository struct {
	db map[string]struct{}
}

func NewTagRepository() *tagRepository {
	return &tagRepository{db: map[string]struct{}{}}
}
