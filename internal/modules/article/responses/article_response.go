package responses

import (
	"fmt"
	articleModels "resedist/internal/modules/article/models"
	userResponse "resedist/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      userResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModels.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      userResponse.ToUser(article.User),
	}
}

func ToArticles(articles []articleModels.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
