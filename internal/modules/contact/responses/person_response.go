package responses

import (
	"fmt"
	userModels "resedist/internal/modules/user/models"
)

type Person struct {
	ID    uint
	Image string
	Name  string
	Email string
}

type People struct {
	Data []Person
}

func ToPerson(user userModels.User) Person {

	return Person{
		ID:    user.ID,
		Name:  user.Name,
		Email: *user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
