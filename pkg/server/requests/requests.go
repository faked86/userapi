package requests

import (
	"net/http"
)

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }
