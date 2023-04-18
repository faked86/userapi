package db

import "userapi/pkg/models"

type UserDB interface {
	GetAll() (UserMap, error)
	GetUser(id string) (models.User, error)
	CreateUser(u models.User) (string, error)
	UpdateUser(id string, displayName string) error
	DeleteUser(id string) error
}
