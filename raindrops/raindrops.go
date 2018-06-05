// Package raindrops converts a number to a string, the contents of which depend
// on the number's factors.
package raindrops

import "fmt"

// factorize returns all the factors of int n.
func factorize(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

/*
Convert converts a number to a raindrop string format.

It follows the following rule:
	- If the number has 3 as a factor, output 'Pling'.
	- If the number has 5 as a factor, output 'Plang'.
	- If the number has 7 as a factor, output 'Plong'.
	- If the number does not have 3, 5, or 7 as a factor,
		just pass the number's digits straight through.
*/
func Convert(n int) string {
	sound := ""
	for _, f := range factorize(n) {
		switch f {
		case 3:
			sound += "Pling"
		case 5:
			sound += "Plang"
		case 7:
			sound += "Plong"
		}
	}

	if sound != "" {
		return sound
	}
	return fmt.Sprintf("%v", n)
}
