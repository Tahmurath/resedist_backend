package repositories

import (
	ContactModels "resedist/internal/modules/contact/models"
)

type PersonRepositoryInterface interface {
	List(limit int) []ContactModels.Person
	Find(id int) ContactModels.Person
	Create(article ContactModels.Person) ContactModels.Person
}
