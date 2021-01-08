package users

import (
	"github.com/nikola-curcic/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	ID 			int64	`json:"id,string"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email		string	`json:"email"`
	DateCreated string	`json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}