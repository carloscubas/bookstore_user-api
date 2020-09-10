package users

import (
	"fmt"
	"github.com/carloscubas/bookstore_user-api/utils/erros"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *erros.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return erros.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *erros.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return erros.NewBadRequestError(fmt.Sprintf("email %s alread registered", user.Email))
		}
		return erros.NewBadRequestError(fmt.Sprintf("user %d alread exist", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
