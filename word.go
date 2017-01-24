package main

import (
	"sort"
	"strings"
)

type Word struct {
	ID                  int
	Word                string
	AlphabetizedVersion string
	Language            string
}

func NewWord(id int, word, alphabetizedVersion, language string) *Word {
	return &Word{
		id,
		word,
		alphabetizedVersion,
		language,
	}
}

func alphabetize(word string) string {
	retVal := strings.Split(word, "")
	sort.Strings(retVal)
	return strings.Join(retVal, "")
}
