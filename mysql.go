package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {
	connection, err := sql.Open("mysql", user+":"+pass+"@tcp(127.0.0.1:3306)/"+database)
	if err != nil {
		return nil, err
	}

	return connection, connection.Ping()
}

func SearchDBForAnagrams(word string) ([]byte, error) {
	rows, err := globalMySQLDB.Query(
		`
		SELECT * FROM words
		WHERE word != ?
		AND alphabetized_version = ?
		`,
		word,
		alphabetize(word),
	)
	if err != nil {
		return nil, err
	}

	words := []Word{}
	for rows.Next() {
		word := Word{}
		err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.AlphabetizedVersion,
			&word.Language,
		)

		if err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	jsonResponse, err := json.MarshalIndent(words, "", " ")
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil

}

func SearchDBForLanguageSpecificAnagrams(word, language string) ([]Word, error) {
	rows, err := globalMySQLDB.Query(
		`
		SELECT * FROM words
		WHERE alphabetized_version = ? 
		AND language = ?
		)
		`,
		alphabetize(word),
		language,
	)
	if err != nil {
		return nil, err
	}

	words := []Word{}
	for rows.Next() {
		word := Word{}
		err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.AlphabetizedVersion,
			&word.Language,
		)

		if err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	return words, nil

}

/*
func PopulateDatabase() {
	InsertFileContentsIntoDB("wordsEn.txt", "english")
	InsertFileContentsIntoDB("wordsEn.txt", "english")
}
*/
