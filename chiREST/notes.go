package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type notesResource struct{}

var notes []Note

type Note struct {
	ID		int
	Title	string
	Author	Author
}

type Author struct {
	Name	string
	Notes	[]int
}
func (rs notesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)
	r.Post("/", rs.Create)
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {

		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
		r.Get("/sync", rs.Sync)
	})

	return r
}

func (rs notesResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos list of stuff.."))
}

func (rs notesResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos create"))
}

func (rs notesResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo get"))
}

func (rs notesResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo update"))
}

func (rs notesResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo delete"))
}

func (rs notesResource) Sync(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo sync"))
}
