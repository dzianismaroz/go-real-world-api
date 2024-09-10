package msg

type ArticlesList struct {
	Articles      []ArticleBody `json:"articles"`
	ArticlesCount uint          `json:"articlesCount"`
}
