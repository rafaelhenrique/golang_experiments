package luhn

import (
	"unicode"
)

// Valid given a creditCardNumber calculate Luhn algorithm and returns if this is a valid number or not
func Valid(creditCardNumber string) (isValid bool) {

	// fmt.Println("=> start: credit card number", creditCardNumber)

	// validate only one number
	if len(creditCardNumber) == 1 {
		return false
	}

	var sum int
	counter := 0
	for i := len(creditCardNumber) - 1; i >= 0; i-- {
		// every first digit from right to left
		r := rune(creditCardNumber[i])
		if unicode.IsDigit(r) {
			// fmt.Println("F number", string(rune(creditCardNumber[i])))
			sum += int(creditCardNumber[i] - '0')
			counter++
		} else if unicode.IsSpace(r) {
			// fmt.Println("F space")
			counter++
		} else {
			return false
		}

		if i-1 < 0 {
			continue
		}

		// every second digit from right to left
		r = rune(creditCardNumber[i-1])
		if unicode.IsDigit(r) {
			if counter%3 == 0 {
				continue
			}
			number := int(creditCardNumber[i-1] - '0')
			number *= 2
			if number > 9 {
				number -= 9
			}
			sum += number
			// fmt.Println("S number", number)
			i--
			counter++
		} else if unicode.IsSpace(r) {
			// fmt.Println("S space")
			counter++
		} else {
			return false
		}
		// fmt.Println("-> Counter", counter)
	}

	// validate if credit card is " 0"
	if unicode.IsSpace(rune(creditCardNumber[0])) && sum == 0 {
		return false
	}

	// make operation to check if sum is divisible by 10
	if sum%10 == 0 {
		// fmt.Println("=> end: sum", sum)
		return true
	}

	// fmt.Println("=> end: sum", sum)
	return false
}
