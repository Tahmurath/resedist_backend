package responses

import (
	"fmt"
	"net/url"
	userModels "resedist/internal/modules/user/models"
)

type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModels.User) User {
	var email string
	if user.Email != nil {
		email = *user.Email
	}
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", url.QueryEscape(user.Name)),
	}
}

func ToUsers(users []userModels.User, expand bool) Users {
	var responseUsers []User

	for _, user := range users {
		responseUsers = append(responseUsers, ToUser(user))
	}

	return Users{
		Data: responseUsers,
	}
}

//response := make([]Department, len(departments))
//
//for i, department := range departments {
//	response[i] = ToDepartment(department, expand)
//}
//
//return Departments{Data: response}
