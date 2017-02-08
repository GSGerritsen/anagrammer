package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleNewSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	email := u.Email
	password := u.Password
	fmt.Printf("--> Received a POST to /login with %s and %s\n", email, password)
}

func HandleNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	email := u.Email
	password := u.Password
	fmt.Printf("--> Received a POST to /signup with %s and %s\n", email, password)
}
