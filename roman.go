package roman
// This is a port of the CPAN-package Roman, V1.23.
// The original package was under the Artistic License.

// Copyright (c) 2012 Stefan Schroeder. All rights reserved. 
// LICENSE is located here: http://golang.org/LICENSE

import "regexp"

// Checks if the string is a valid Roman number.
// Checks against a hilariuosly complex regex.
func IsRoman(arg string)(bool) {
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
func Roman(arg int)(string) {

    figure := []int{1000,100,10,1}
    roman_digitA := []string{
  1 : "I",
	10 : "X",
	100 : "C",
	1000 : "M",
    }
    roman_digitB := []string{
	1 : "V",
	10 : "L",
	100 : "D",
	1000 : "MMMMM",
    }

    if arg < 0 || arg > 4000 {
	return "ROMAN_OUT_OF_RANGE"
    }

    roman_slice := []string{}
    x := ""

    for _, f := range figure {
        digit, i, v := int(arg / f), roman_digitA[f], roman_digitB[f]
        switch digit {
        case 1: roman_slice = append(roman_slice, string(i) )
        case 2: roman_slice = append(roman_slice, string(i) + string(i))
        case 3: roman_slice = append(roman_slice, string(i) + string(i) + string(i) )
        case 4: roman_slice = append(roman_slice, string(i) + string(v) )
        case 5: roman_slice = append(roman_slice, string(v) )
        case 6: roman_slice = append(roman_slice, string(v) + string(i) )
        case 7: roman_slice = append(roman_slice, string(v) + string(i)+ string(i) )
        case 8: roman_slice = append(roman_slice, string(v) + string(i)+ string(i)+ string(i) )
        case 9: roman_slice = append(roman_slice, string(i) + string(x)  )
	}

        arg -= digit * f
        x = i
    }
    ret := ""
    for _, e := range roman_slice {
        ret += e
    }
    return ret
}

// Convert the argument Roman string to an arabic integer.
// Returns -1 if not a valid Roman number.
func Arabic(arg string)(int) {

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

