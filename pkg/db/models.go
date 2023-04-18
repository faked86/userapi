package db

import "userapi/pkg/models"

type UserMap map[string]models.User

type UserStore struct {
	Increment int     `json:"increment"`
	Map       UserMap `json:"list"`
}
