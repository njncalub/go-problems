// Package scrabble calculates the total score for a word.
package scrabble

import "strings"

/*
Score computes the scrabble score for a specific word.
*/
func Score(word string) int {
	total, word := 0, strings.ToUpper(strings.TrimSpace(word))

	for _, c := range word {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			total++
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		}
	}

	return total
}
