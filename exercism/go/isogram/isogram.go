package isogram

import (
	"regexp"
	"strings"
	"unicode"
)

var regexPattern, _ = regexp.Compile("[^a-zA-Z]+")

// IsIsogramV1 receive a word and return true if word is a isogram and false if it is not
func IsIsogramV1(word string) (isIsogram bool) {

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

// IsIsogramV2 receive a word and return true if word is a isogram and false if it is not
func IsIsogramV2(word string) (isIsogram bool) {
	cleanedWord := strings.ToLower(word)
	cleanedWord = strings.ReplaceAll(cleanedWord, "-", "")
	cleanedWord = strings.ReplaceAll(cleanedWord, " ", "")

	for _, rune := range cleanedWord {
		count := strings.Count(cleanedWord, string(rune))
		if count > 1 {
			return false
		}
	}
	return true
}

// IsIsogram receive a word and return true if word is a isogram and false if it is not
func IsIsogram(word string) (isIsogram bool) {
	for _, rune := range word {
		if unicode.IsLetter(rune) == false {
			continue
		}
		lowerRune := unicode.ToLower(rune)
		upperRune := unicode.ToUpper(rune)
		count := strings.Count(word, string(lowerRune))
		count += strings.Count(word, string(upperRune))
		if count > 1 {
			return false
		}
	}
	return true
}
