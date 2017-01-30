package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var globalMySQLDB *sql.DB

func init() {
	db, err := ConnectToMySQL()
	if err != nil {
		panic(err)
	}

	globalMySQLDB = db
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	return router
}

func main() {
	//InsertFileContentsIntoDB("francais.txt", "french")
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/search/:word", HandleHomeGet)
	router.Handle("POST", "/", HandleWordSearch)

	// Testing routes for Vue stuff

	router.Handle("GET", "/vue-tests", HandleVueTests)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
