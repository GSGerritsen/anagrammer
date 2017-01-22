package main

import (
	"fmt"
)

// Just printing to console for now to check that things are working
func PrintAnagrams(word string) {
	words, err := SearchDBForAnagrams(word)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Anagrams for %s:\n", word)
	for i := 0; i < len(words); i++ {
		fmt.Printf("%s\n", words[i].word)
	}
}
