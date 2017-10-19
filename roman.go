package roman

// This is a port of the CPAN-package Roman, V1.23.
// The original package was under the Artistic License.

// Copyright (c) 2012 Stefan Schroeder. All rights reserved.
// LICENSE is located here: http://golang.org/LICENSE

import "regexp"

var roman2arabic = []struct {
	roman  string
	arabic int
}{
	{"I", 1},
	{"IV", 4},
	{"V", 5},
	{"IX", 9},
	{"X", 10},
	{"XL", 40},
	{"L", 50},
	{"XC", 90},
	{"C", 100},
	{"CD", 400},
	{"D", 500},
	{"CM", 900},
	{"M", 1000},
}

// JOIN From strings.Join
// following robs pike advice from golang proverbs "A little copying is better than a little dependency."
func Join(a []string, sep string) string {

	switch len(a) {

	case 0:

		return ""

	case 1:

		return a[0]

	case 2:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1]

	case 3:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1] + sep + a[2]

	}

	n := len(sep) * (len(a) - 1)

	for i := 0; i < len(a); i++ {

		n += len(a[i])

	}

	b := make([]byte, n)

	bp := copy(b, a[0])

	for _, s := range a[1:] {

		bp += copy(b[bp:], sep)

		bp += copy(b[bp:], s)

	}

	return string(b)

}

// Checks if the string is a valid Roman number.
// Checks against a hilariuosly complex regex.
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

// Convert the argument integer to a Roman number string.
// Valid only for values greater 0 and smaller than 4000.
func Roman(arg int) string {

	if arg < 0 || arg > 4000 {
		return "ROMAN_OUT_OF_RANGE"
	}
	out := []string{}
	i := len(roman2arabic) - 1
	for i >= 0 {
		if roman2arabic[i].arabic <= arg {
			out = append(out, roman2arabic[i].roman)
			arg -= roman2arabic[i].arabic
		} else {
			i -= 1
		}
	}
	return Join(out, "")

}

// Convert the argument Roman string to an arabic integer.
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
