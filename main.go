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
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("POST", "/", HandleWordSearch)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
