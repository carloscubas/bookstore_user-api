package users

import (
	"fmt"
	"github.com/carloscubas/bookstore_user-api/datasources/mysql/users_db"
	"github.com/carloscubas/bookstore_user-api/utils/date_utils"
	"github.com/carloscubas/bookstore_user-api/utils/erros"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *erros.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
