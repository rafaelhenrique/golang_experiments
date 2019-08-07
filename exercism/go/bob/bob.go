package bob

import (
	"strings"
	"unicode"
)

// isUpper receive s and return true when all letter runes is upper
func isUpper(s string) bool {
	lettersCount := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			lettersCount++
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	if lettersCount == 0 {
		return false
	}
	return true
}

// isQuestion receive s and return true when s have a interrogation sign on last rune
func isQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	lastElement := len(s) - 1
	lastRune := s[lastElement]
	if lastRune == '?' {
		return true
	}
	return false
}

// Hey receive some remark and return some answer of bob
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	} else if isQuestion(remark) && isUpper(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if isUpper(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
