package repositories

import (
	ContactModels "resedist/internal/modules/contact/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func New() *PersonRepository {
	return &PersonRepository{
		DB: database.Connection(),
	}
}

func (PersonRepository *PersonRepository) List(limit int) []ContactModels.Person {
	var people []ContactModels.Person

	PersonRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&people)

	return people
}

func (PersonRepository *PersonRepository) Find(id int) ContactModels.Person {
	var person ContactModels.Person

	PersonRepository.DB.Joins("User").First(&person, id)

	return person
}

func (PersonRepository *PersonRepository) Create(person ContactModels.Person) ContactModels.Person {
	var newPerson ContactModels.Person

	PersonRepository.DB.Create(&person).Scan(&newPerson)

	return newPerson
}
