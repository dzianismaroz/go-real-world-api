package msg

type CreateArticleMessage struct {
	Content ArticleContent `json:"article"`
}

type ArticleContent struct {
	Body         string   `json:"body"`
	Descrtiption string   `json:"description"`
	Tags         []string `json:"tagList"`
	Title        string   `json:"title"`
}

func (a *CreateArticleMessage) IsValid() bool {
	switch {
	case a.Content.Title == "":
		return false
	case a.Content.Body == "":
		return false
	default:
		return true
	}
}
