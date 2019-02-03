package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/user/(name}/{email}", NewUser).Methods("POST")
	r.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Go ORM API")
	handleRequests()
}
