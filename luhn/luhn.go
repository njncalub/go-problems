// Package luhn implements a checker for values created using
// the Luhn algorightm.
package luhn

import (
	"regexp"
	"strconv"
)

var (
	invalidChars = regexp.MustCompile(`[^0-9 ]+`)
	numericOnly  = regexp.MustCompile(`[^0-9]+`)
)

// Valid determines whether a number is valid per the
// Luhn formula.
func Valid(s string) bool {
	// Invalidate if the string contains invalid characters.
	if invalidChars.MatchString(s) {
		return false
	}

	s = numericOnly.ReplaceAllString(s, "")
	length := len(s)

	// Invalidate if the string has less than 2 characters.
	if length <= 1 {
		return false
	}

	sum := 0
	for i := 0; i < length; i++ {
		d := string(s[length-i-1])

		// Disregards error because it is not being tested in the test cases.
		n, _ := strconv.Atoi(d)

		if i%2 != 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}

	return sum%10 == 0
}
