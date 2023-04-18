package db

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strconv"

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

func (db *UserDB) GetAll() (UserMap, error) {
	s, err := db.readDB()

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

func (db *UserDB) CreateUser(u models.User) (string, error) {
	s, err := db.readDB()
	if err != nil {
		return "", err
	}

	s.Increment++
	id := strconv.Itoa(s.Increment)
	s.Map[id] = u

	if err := db.writeDB(s); err != nil {
		return "", err
	}

	return id, nil
}

func (db *UserDB) readDB() (UserStore, error) {
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

func (db *UserDB) writeDB(s UserStore) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(db.filename, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
