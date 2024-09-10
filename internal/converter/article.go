package converter

import (
	"rwa/pkg/model"
	"rwa/pkg/msg"
	"time"
)

func ToArticleEntity(articleReq *msg.CreateArticleMessage, authorId uint64) model.Article {
	return model.Article{
		AuthorId:    authorId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: articleReq.Content.Descrtiption,
		TagList:     articleReq.Content.Tags,
		Title:       articleReq.Content.Title,
		Body:        articleReq.Content.Body,
	}
}

func ToArticleMsg(article model.Article, author msg.UserProfile) msg.Article {
	return msg.Article{
		Content: msg.ArticleBody{
			Author: msg.InnerContentShort{
				Image:     author.Inner.Image,
				Bio:       author.Inner.Bio,
				Following: author.Inner.Following,
				Username:  author.Inner.Username,
			},
			Body:            article.Body,
			CreatedAt:       article.CreatedAt,
			Description:     article.Description,
			Favourited:      article.Favorited,
			FavouritedCount: article.FavouritesCount,
			Slug:            article.Slug,
			Tags:            article.TagList,
			Title:           article.Title,
			UpdatedAt:       article.UpdatedAt,
		},
	}
}

func ToArticleBodyMsg(article model.Article, author msg.UserProfile) msg.ArticleBody {
	return msg.ArticleBody{
		Author: msg.InnerContentShort{
			Image:     author.Inner.Image,
			Bio:       author.Inner.Bio,
			Following: author.Inner.Following,
			Username:  author.Inner.Username,
		},
		Body:            article.Body,
		CreatedAt:       article.CreatedAt,
		Description:     article.Description,
		Favourited:      article.Favorited,
		FavouritedCount: article.FavouritesCount,
		Slug:            article.Slug,
		Tags:            article.TagList,
		Title:           article.Title,
		UpdatedAt:       article.UpdatedAt,
	}
}

func ToArticlesList(articles []msg.ArticleBody) msg.ArticlesList {
	return msg.ArticlesList{
		Articles:      articles,
		ArticlesCount: uint(len(articles)),
	}
}
