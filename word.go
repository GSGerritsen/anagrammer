package main

import (
	"sort"
	"strings"
)

type Word struct {
	id                   int
	word                 string
	alphabetized_version string
	language             string
}

func NewWord(id int, word, alphabetized_version, language string) *Word {
	return &Word{
		id,
		word,
		alphabetized_version,
		language,
	}
}

func alphabetize(word string) string {
	retVal := strings.Split(word, "")
	sort.Strings(retVal)
	return strings.Join(retVal, "")
}
