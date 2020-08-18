package services

import (
	"github.com/sug5806/bookstore_users-api/domain/users"
	"github.com/sug5806/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil

}

func GetUser() {}

func FindUser() {}
