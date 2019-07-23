package isogram

import (
	"regexp"
	"strings"
)

var regexPattern, _ = regexp.Compile("[^a-zA-Z]+")

// IsIsogram receive a word and return true if word is a isogram and false if it is not
func IsIsogram(word string) (isIsogram bool) {

	cleanedWord := regexPattern.ReplaceAllString(word, "")
	cleanedWord = strings.ToLower(cleanedWord)

	for firstIndex, firstRune := range cleanedWord {
		for secondIndex, secondRune := range cleanedWord {
			if firstRune == secondRune && firstIndex != secondIndex {
				return false
			}
		}
	}
	return true
}
