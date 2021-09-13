package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbc := createMongoDBClient("<ATLAS_URI_HERE>")
	defer dbc.Close()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/events", func(r chi.Router) {
		r.Put("/", dbc.create)
		r.Get("/", dbc.read)
		r.Patch("/", dbc.update)
		r.Delete("/", dbc.delete)
	})

	http.ListenAndServe(":3000", r)
}
