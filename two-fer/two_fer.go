// Package twofer returns a Two-fer, which is short for two for one: One for
// you and one for me.
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns the formatted two-fer string. If no name is given, the
// name defaults to "you".
func ShareWith(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
