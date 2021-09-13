package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/portobello-boy/MicroservicesDemo/server"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbc := server.CreateMongoDBClient("mongodb://localhost:27017")
	defer dbc.Close()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/events", func(r chi.Router) {
		r.Put("/", dbc.Create)
		r.Get("/", dbc.Read)
		r.Patch("/", dbc.Update)
		r.Delete("/", dbc.Delete)
	})

	http.ListenAndServe(":3000", r)
}
