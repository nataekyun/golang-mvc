package service

import (
	"MVC/domain"
)

func GetUser(userID int64) (*domain.User, error) {

	return domain.GetUser(userID)

}
