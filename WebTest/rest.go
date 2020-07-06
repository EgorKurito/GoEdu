package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests()  {
	myRouter := mux.NewRouter().StrictSlash(true)

	//http.HandleFunc("/", homePage)
	myRouter.HandleFunc("/", homePage)

	//http.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "Key: " + key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Id		string `json:"Id"`
	Title	string `json:"Title"`
	Desc	string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main()  {
	Articles = []Article{
		Article{
			Id:			"1",
			Title:		"Hello",
			Desc:		"Article Description",
			Content:	"Article Content",
		},
		Article{
			Id:			"2",
			Title:		"Hello 2",
			Desc:		"Article Description",
			Content:	"Article Content",
		},
	}
	handleRequests()

}
