package db

import (
	"encoding/json"
	"io/ioutil"

	"userapi/pkg/models"
	"userapi/pkg/server/customerrors"
)

type UserDB struct {
	filename string
}

func NewUserDB(filename string) UserDB {
	return UserDB{
		filename: filename,
	}
}

func (db *UserDB) ReadDB() (UserStore, error) {
	f, err := ioutil.ReadFile(db.filename)
	if err != nil {
		return UserStore{}, err
	}

	s := UserStore{}

	if err = json.Unmarshal(f, &s); err != nil {
		return UserStore{}, err
	}

	return s, nil
}

func (db *UserDB) GetAll() (UserMap, error) {
	s, err := db.ReadDB()

	return s.Map, err
}

func (db *UserDB) GetUser(id string) (models.User, error) {
	usrs, err := db.GetAll()
	if err != nil {
		return models.User{}, err
	}

	if u, ok := usrs[id]; ok {
		return u, nil
	}

	return models.User{}, customerrors.ErrUserNotFound
}
