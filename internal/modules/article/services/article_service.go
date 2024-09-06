package services

import (
	"errors"
	ArticelRepository "resedist/internal/modules/article/repositories"
	ArticleResponse "resedist/internal/modules/article/responses"
)

type ArticleService struct {
	articelRepository ArticelRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articelRepository: ArticelRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {

	articles := ArticleService.articelRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (ArticleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {

	articles := ArticleService.articelRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (ArticleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {

	var response ArticleResponse.Article
	article := ArticleService.articelRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponse.ToArticle(article), nil
}
