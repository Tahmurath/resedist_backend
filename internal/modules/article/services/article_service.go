package services

import (
	"errors"
	ArticleModel "resedist/internal/modules/article/models"
	ArticelRepository "resedist/internal/modules/article/repositories"
	"resedist/internal/modules/article/requests/articles"
	ArticleResponse "resedist/internal/modules/article/responses"
	UserResponse "resedist/internal/modules/user/responses"
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

func (ArticleService *ArticleService) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := ArticleService.articelRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error in creating article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}
