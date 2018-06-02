// Package acronym converts a phrase to its acronym.
package acronym

import (
	"strings"
	"unicode"
)

// checkSeparator takes a rune character and checks if it should be used as a
// separator when splitting a string.
func checkSeparator(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

// Abbreviate takes a string and returns it in uppercased acronym format.
func Abbreviate(s string) string {
	acronym := ""
	for _, word := range strings.FieldsFunc(s, checkSeparator) {
		acronym += string(word[0])
	}
	return strings.ToUpper(acronym)
}
