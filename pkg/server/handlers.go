package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"userapi/pkg/models"
	ce "userapi/pkg/server/customerrors"
	"userapi/pkg/server/requests"
	resp "userapi/pkg/server/responses"
)

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usrs, err := s.db.GetAll()
	if err != nil {
		resp.RenderInternalError(w, r, err)
		return
	}
	render.JSON(w, r, usrs)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := s.db.GetUser(id)

	if err == ce.ErrUserNotFound {
		resp.RenderNotFound(w, r, err)
		return
	}

	if err != nil {
		resp.RenderInternalError(w, r, err)
		return
	}

	render.JSON(w, r, u)
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		resp.RenderInvalidRequest(w, r, err)
		return
	}

	if request.DisplayName == "" {
		resp.RenderInvalidRequest(w, r, ce.ErrEmptyName)
		return
	}

	if request.Email == "" {
		resp.RenderInvalidRequest(w, r, ce.ErrEmptyEmail)
		return
	}

	u := models.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id, err := server.db.CreateUser(u)
	if err != nil {
		resp.RenderInternalError(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		resp.RenderInvalidRequest(w, r, err)
		return
	}

	if request.DisplayName == "" {
		resp.RenderInvalidRequest(w, r, ce.ErrEmptyName)
		return
	}

	id := chi.URLParam(r, "id")

	err := server.db.UpdateUser(id, request.DisplayName)

	if err == ce.ErrUserNotFound {
		resp.RenderNotFound(w, r, err)
		return
	}

	if err != nil {
		resp.RenderInternalError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := server.db.DeleteUser(id)

	if err == ce.ErrUserNotFound {
		resp.RenderNotFound(w, r, err)
		return
	}

	if err != nil {
		resp.RenderInternalError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
}
