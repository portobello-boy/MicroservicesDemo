package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/portobello-boy/MicroservicesDemo/CRUD/server"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbc := server.CreateMongoDBClient("mongodb://localhost:27017")
	defer dbc.Close()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	r.Route("/events", func(r chi.Router) {
		r.Put("/", dbc.Create)
		r.Get("/", dbc.ReadAll)
		r.Get("/{id}", dbc.Read)
		r.Patch("/", dbc.Update)
		r.Delete("/{id}", dbc.Delete)
	})

	log.Print("Server listening on port 3000...")
	http.ListenAndServe(":3000", r)
}
