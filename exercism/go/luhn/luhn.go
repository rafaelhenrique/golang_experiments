package luhn

import (
	"strings"
	"unicode"
)

// Valid given a creditCardNumber calculate Luhn algorithm and returns if this is a valid number or not
func Valid(creditCardNumber string) (isValid bool) {

	// validate and sanitize sequence of numbers
	var sb strings.Builder
	sb.Grow(len(creditCardNumber))
	for _, r := range creditCardNumber {
		if unicode.IsNumber(r) {
			sb.WriteRune(r)
		} else if unicode.IsSpace(r) {
			continue
		} else {
			return false
		}
	}
	sanitizedCardNumber := sb.String()

	// validate only one number
	if len(sanitizedCardNumber) == 1 {
		return false
	}

	var sum, fNumber, sNumber int
	for first, second := len(sanitizedCardNumber)-1, len(sanitizedCardNumber)-2; first >= 0; first, second = first-2, second-2 {

		// logic to every second number from right to left
		if second >= 0 {
			sNumber = int(sanitizedCardNumber[second] - '0')
			sNumber *= 2
			if sNumber > 9 {
				sNumber -= 9
			}
			sum += sNumber
		}

		fNumber = int(sanitizedCardNumber[first] - '0')
		sum += fNumber
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
