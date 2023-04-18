package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"userapi/pkg/db"
)

type Server struct {
	port   string
	db     *db.UserDB
	router *chi.Mux
}

func NewServer(port string, db *db.UserDB) Server {
	return Server{
		port:   port,
		db:     db,
		router: chi.NewRouter(),
	}
}

func (s *Server) Start() {
	s.configureRouter()

	log.Printf("HTTP server starting at port: %s", s.port)

	err := http.ListenAndServe(":"+s.port, s.router)
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func (s *Server) configureRouter() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", s.GetAllUsers)
				r.Post("/", s.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", s.GetUser)
					r.Patch("/", s.UpdateUser)
					r.Delete("/", s.DeleteUser)
				})
			})
		})
	})
}
