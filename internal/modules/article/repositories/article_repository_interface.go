package repositories

import (
	ArticleModels "resedist/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModels.Article
}
