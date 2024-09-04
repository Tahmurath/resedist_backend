package repositories

import (
	"gorm.io/gorm"
	ArticleModels "resedist/internal/modules/article/models"
	"resedist/pkg/database"
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
