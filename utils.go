package main

import (
	"strings"
	"unicode"
)

// titleCase returns a copy of the string s with the first letter of each word capitalized.
// This is a replacement for the deprecated strings.Title function.
func titleCase(s string) string {
	// Split the string into words
	words := strings.Fields(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return ' '
	}, s))

	// Capitalize the first letter of each word
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}

	// Join the words back together
	return strings.Join(words, " ")
}

// isPNGFile checks if the given file path has a .png extension
func isPNGFile(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".png")
}

// isGoFile checks if the given file path has a .go extension
func isGoFile(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".go")
}
