package server

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"userapi/pkg/db"
	"userapi/pkg/models"
	ce "userapi/pkg/server/customerrors"
	"userapi/pkg/server/requests"
	resp "userapi/pkg/server/responses"
)

const store = `users.json`

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

	server.db.UpdateUser(id, request.DisplayName)

	render.Status(r, http.StatusNoContent)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := db.UserStore{}
	_ = json.Unmarshal(f, &s)

	id := chi.URLParam(r, "id")

	if _, ok := s.Map[id]; !ok {
		resp.RenderInvalidRequest(w, r, ce.ErrUserNotFound)
		return
	}

	delete(s.Map, id)

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}
