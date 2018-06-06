// Package reverse reverses a string.
package reverse

// String returns the reversed string.
func String(s string) string {
	// Manually deconstruct the string to a slice of runes
	// to properly identify mixed characters as 1 rune.
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
