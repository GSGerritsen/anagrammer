package main

import (
	"bufio"
	"fmt"
	"os"
)

// Just printing to console for now to check that things are working
func PrintAnagrams(word string) {
	words, err := SearchDBForAnagrams(word)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(words))
}

func SearchOnInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a word to find its anagrams! ")
	for scanner.Scan() {
		PrintAnagrams(scanner.Text())
	}

}
