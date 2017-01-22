package main

import (
	"database/sql"
)

var globalMySQLDB *sql.DB

func init() {
	db, err := ConnectToMySQL()
	if err != nil {
		panic(err)
	}

	globalMySQLDB = db
}

func main() {
	PrintAnagrams("reams")

}
