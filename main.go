package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"userapi/methods"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", methods.SearchUsers)
				r.Post("/", methods.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", methods.GetUser)
					r.Patch("/", methods.UpdateUser)
					r.Delete("/", methods.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
