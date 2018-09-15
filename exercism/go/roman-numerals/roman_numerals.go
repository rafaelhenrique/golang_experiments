package romannumerals

import (
	"errors"
)

func ToRomanNumeral(arabicNumber int) (romanNumber string, err error) {
	if arabicNumber > 3000 {
		return "", errors.New("ToRomanNumeral: cannot convert arabic numbers greater than 3000")
	}

	if arabicNumber <= 0 {
		return "", errors.New("ToRomanNumeral: cannot convert arabic numbers lower than equal 0")
	}

	unitNumbers := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
	}
	decimalNumbers := map[int]string{
		1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC",
	}
	hundredNumbers := map[int]string{
		1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM",
	}
	thousandNumbers := map[int]string{
		1: "M", 2: "MM", 3: "MMM",
	}

	thousandPart := arabicNumber / 1000
	hundredPart := arabicNumber / 100
	decimalPart := arabicNumber / 10
	unitPart := arabicNumber % 10

	if thousandPart != 0 {
		romanNumber += thousandNumbers[thousandPart]
	}

	if hundredPart != 0 {
		romanNumber += hundredNumbers[hundredPart]
		decimalPart = (arabicNumber - (100 * hundredPart)) / 10
	}

	if decimalPart != 0 {
		romanNumber += decimalNumbers[decimalPart]
	}

	if unitPart != 0 {
		romanNumber += unitNumbers[unitPart]
	}

	return romanNumber, nil
}
