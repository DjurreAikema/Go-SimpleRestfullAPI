package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, req *http.Request) {
	articles := Articles{
		{Title: "Test title", Desc: "Test description", Content: "Test content"},
	}

	fmt.Println("Endoint hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequest()
}
