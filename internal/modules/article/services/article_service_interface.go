package services

import (
	"resedist/internal/modules/article/requests/articles"
	ArticleResponse "resedist/internal/modules/article/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
