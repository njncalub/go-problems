/*
Package isogram determines if a word or phrase is an isogram.

An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

	- lumberjacks
	- background
	- downstream
	- six-year-old

The word *isograms*, however, is not an isogram, because the s repeats.
*/
package isogram

import "strings"

func removeInvalidChars(r rune) rune {
	if strings.ContainsRune(" -", r) {
		return -1
	}
	return r
}

// IsIsogram checks if the given word is an isogram.
func IsIsogram(s string) bool {
	s = strings.Map(removeInvalidChars, strings.ToLower(s))
	m := map[rune]int{}

	for _, c := range s {
		if _, found := m[c]; found {
			return false
		}
		m[c] = 1
	}

	return true
}
