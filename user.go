package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Endpoint: all users")
}

func NewUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Endpoint: new user")
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Endpoint: delete user")
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Endpoint: update user")
}
