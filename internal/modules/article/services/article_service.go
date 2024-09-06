package services

import (
	articleModels "resedist/internal/modules/article/models"
	ArticelRepository "resedist/internal/modules/article/repositories"
)

type ArticleService struct {
	articelRepository ArticelRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articelRepository: ArticelRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() []articleModels.Article {

	return ArticleService.articelRepository.List(4)
}

func (ArticleService *ArticleService) GetStoriesArticles() []articleModels.Article {

	return ArticleService.articelRepository.List(6)
}
