package seeder

import (
	"fmt"
	"log"
	articleModels "resedist/internal/modules/article/models"
	userModels "resedist/internal/modules/user/models"
)

func ArticleSeed(user userModels.User) articleModels.Article {

	var article articleModels.Article

	for i := 1; i <= 10; i++ {
		article = articleModels.Article{Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Content %d", i), UserID: user.ID}
		db.Create(&article) // pass pointer of data to Create

		log.Printf("Article created with title: %s", article.Title)
	}

	return article
}
