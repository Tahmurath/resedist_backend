package repositories

import (
	ArticleModels "resedist/internal/modules/article/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (ArticleRepository *ArticleRepository) List(limit int) []ArticleModels.Article {
	var articles []ArticleModels.Article

	ArticleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (ArticleRepository *ArticleRepository) Find(id int) ArticleModels.Article {
	var article ArticleModels.Article

	ArticleRepository.DB.Joins("User").First(&article, id)

	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModels.Article) ArticleModels.Article {
	var newArticle ArticleModels.Article

	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}
