package services

import (
	articleModels "resedist/internal/modules/article/models"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() []articleModels.Article
	GetStoriesArticles() []articleModels.Article
}
