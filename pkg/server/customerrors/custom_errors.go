package customerrors

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmptyName    = errors.New("display_name field is empty")
	ErrEmptyEmail   = errors.New("email field is empty")
)
