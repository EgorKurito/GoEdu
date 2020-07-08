package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Handler func(writer http.ResponseWriter, request *http.Request) error

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if err := h(writer, request); err != nil {
		writer.WriteHeader(503)
		writer.Write([]byte("Bad"))

	}
}

func homePage(writer http.ResponseWriter, request *http.Request) error {
	q := request.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}
	//fmt.Fprintf(writer, "Welcome to the HomePage!")
	_, _ = writer.Write([]byte("Welcome to the HomePage!"))
	fmt.Println("Endpoint Hit: homePage")
	return nil
}

func handleRequest()  {
	myRouter := chi.NewRouter()
	myRouter.Use(middleware.RequestID)
	myRouter.Use(middleware.Logger)
	myRouter.Use(middleware.Recoverer)

	myRouter.Method("GET", "/", Handler(homePage))

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main()  {
	handleRequest()
}