package roman

// This is a port of the CPAN-package Roman, V1.23.
// The original package was under the Artistic License.

// Copyright (c) 2012 Stefan Schroeder. All rights reserved.
// LICENSE is located here: http://golang.org/LICENSE

import (
	"errors"
	"strings"
)

// IsRoman checks if the string is a valid Roman number.
// Checks against a hilariuosly complex regex.
/* commented out until further notice because it doesnt work.
func IsRoman(arg string) bool {
	if arg == "" {
		return false
	}
	c := []byte(arg)
	if check, _ := regexp.Match("I{0,3}", c); check == true {
		//if check, _ := regexp.Match("(?:M{0,3})(?:D?C{0,3}|C[DM])(?:L?X{0,3}|X[LC])(?:V?I{0,3}|I[VX])", c); check == true {
		return true
	}
	return false
}
*/


var (
	romanFigure = []int{1000, 100, 10, 1}
	romanDigitA = []string{1: "I", 10: "X", 100: "C", 1000: "M"}
	romanDigitB = []string{1: "V", 10: "L", 100: "D", 1000: "MMMMM"}

	ErrOutOfRange = errors.New("out of range")
)

// Roman converts the argument integer to a Roman number string.
// Valid only for values greater 0 and smaller than 4000.
func Roman(arg int) (string, error) {
	// Return early in case of invalid argument.
	if arg < 1 || arg > 4000 {
		return "", ErrOutOfRange
	}

	var roman strings.Builder
	x := ""

	for _, f := range romanFigure {
		digit, i, v := int(arg/f), romanDigitA[f], romanDigitB[f]
		switch digit {
		case 1:
			roman.WriteString(i)
		case 2:
			roman.WriteString(i + i)
		case 3:
			roman.WriteString(i + i + i)
		case 4:
			roman.WriteString(i + v)
		case 5:
			roman.WriteString(v)
		case 6:
			roman.WriteString(v + i)
		case 7:
			roman.WriteString(v + i + i)
		case 8:
			roman.WriteString(v + i + i + i)
		case 9:
			roman.WriteString(i + x)
		}

		arg -= digit * f
		x = i
	}

	return roman.String(), nil
}

// Arabic converts the argument Roman string to an arabic integer.
// Returns -1 if not a valid Roman number.
func Arabic(arg string) int {

	roman2arabic := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	/*if ! IsRoman(arg) {
	   return -1
	}*/
	lastDigit := 1000
	arabic := 0
	c := []byte(arg)
	for _, v := range c {
		digit := roman2arabic[string(v)]
		if lastDigit < digit {
			arabic -= 2 * lastDigit
		}
		lastDigit = digit
		arabic += lastDigit
	}
	return arabic
}
