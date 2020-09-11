package users

import (
	"fmt"
	"github.com/carloscubas/bookstore_user-api/datasources/mysql/users_db"
	"github.com/carloscubas/bookstore_user-api/utils/date_utils"
	"github.com/carloscubas/bookstore_user-api/utils/erros"
	"strings"
)

const (
	indexUniqueEmail            = "email_UNIQUE"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() *erros.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return erros.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		fmt.Println(getErr)
		return erros.NewInternalServerError(fmt.Sprintf("error whe trying to get user %d: %s", user.Id, getErr.Error()))
	}

	return nil
}

func (user *User) Save() *erros.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return erros.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return erros.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return erros.NewInternalServerError(
			fmt.Sprintf("error when try to save user %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return erros.NewInternalServerError(fmt.Sprintf("error when try to save user %s", err.Error()))
	}

	user.Id = userId

	return nil
}
