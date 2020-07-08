package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main()  {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("."))
	})

	r.Mount("/notes", notesResource{}.Routes())

	http.ListenAndServe(":3333", r)
}