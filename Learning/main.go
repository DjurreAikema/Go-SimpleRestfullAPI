package Learning

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Endoint hit: test post endpoint")
}

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homePage)
	r.HandleFunc("/articles", allArticles).Methods("GET")
	r.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequest()
}
