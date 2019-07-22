package scrabble

import (
	"unicode"
)

var letterValues = map[rune]int{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10,
}

// ScoreWithMap compute the scrabble score given an word
func ScoreWithMap(word string) (score int) {
	for _, rune := range word {
		letter := unicode.ToLower(rune)
		score += letterValues[letter]
	}
	return
}

// Score compute the scrabble score given an word
func Score(word string) (score int) {
	for _, rune := range word {
		letter := unicode.ToLower(rune)
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++

		case 'd', 'g':
			score += 2

		case 'b', 'c', 'm', 'p':
			score += 3

		case 'f', 'h', 'v', 'w', 'y':
			score += 4

		case 'k':
			score += 5

		case 'j', 'x':
			score += 8

		case 'q', 'z':
			score += 10

		}
	}
	return
}
