package services

import (
	"github.com/sug5806/bookstore_users-api/domain/users"
	"github.com/sug5806/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}

func GetUser() {}

func FindUser() {}
