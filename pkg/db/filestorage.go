package db

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strconv"

	ce "userapi/pkg/customerrors"
	"userapi/pkg/models"
)

type FileDB struct {
	filename string
}

func NewFileDB(filename string) FileDB {
	return FileDB{
		filename: filename,
	}
}

func (db *FileDB) GetAll() (UserMap, error) {
	s, err := db.readDB()

	return s.Map, err
}

func (db *FileDB) GetUser(id string) (models.User, error) {
	usrs, err := db.GetAll()
	if err != nil {
		return models.User{}, err
	}

	if u, ok := usrs[id]; ok {
		return u, nil
	}

	return models.User{}, ce.ErrUserNotFound
}

func (db *FileDB) CreateUser(u models.User) (string, error) {
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

func (db *FileDB) UpdateUser(id string, displayName string) error {
	s, err := db.readDB()
	if err != nil {
		return err
	}

	if _, ok := s.Map[id]; !ok {
		return ce.ErrUserNotFound
	}

	u := s.Map[id]
	u.DisplayName = displayName
	s.Map[id] = u

	if err := db.writeDB(s); err != nil {
		return err
	}

	return nil
}

func (db *FileDB) DeleteUser(id string) error {
	s, err := db.readDB()
	if err != nil {
		return err
	}

	if _, ok := s.Map[id]; !ok {
		return ce.ErrUserNotFound
	}

	delete(s.Map, id)

	if err := db.writeDB(s); err != nil {
		return err
	}

	return nil
}

func (db *FileDB) readDB() (UserStore, error) {
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

func (db *FileDB) writeDB(s UserStore) error {
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
