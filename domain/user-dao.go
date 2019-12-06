package domain

import (
	"fmt"

	"errors"
)

var (
	users = map[int64]*User{
		123: {UserID: 123, FirstName: "Na", LastName: "taekyun", Email: "bbsmax@nate.com"},
	}
)

func GetUser(userID int64) (*User, error) {

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("user [%v] not exist", userID))
}
