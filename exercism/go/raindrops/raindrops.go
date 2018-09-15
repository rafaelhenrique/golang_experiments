package raindrops

import "strconv"

// Convert number to raindropOutput
func Convert(number int) (raindropOutput string) {

	if number%3 == 0 {
		raindropOutput += "Pling"
	}

	if number%5 == 0 {
		raindropOutput += "Plang"
	}

	if number%7 == 0 {
		raindropOutput += "Plong"
	}

	if raindropOutput == "" {
		return strconv.Itoa(number)
	}

	return
}
