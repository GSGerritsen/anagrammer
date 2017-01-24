package main

import (
	"bufio"
	_ "database/sql"
	//"fmt"
	"log"
	"os"
)

// To be used once per dictionary file
func InsertFileContentsIntoDB(filename, lang string) {
	file, err := os.Open("./data/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		word := scanner.Text()
		alphabetizedVersion := alphabetize(word)
		_, err := globalMySQLDB.Exec(
			`
			REPLACE INTO words (
				word,
				alphabetized_version,
				language
				)
			VALUE (
				?,
				?,
				?
			)
			`,
			word,
			alphabetizedVersion,
			lang,
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	if scanErr := scanner.Err(); scanErr != nil {
		log.Fatal(scanErr)
	}

}
