// Package twofer provides a simple way to generate a phrase for sharing an item.
package twofer

import "fmt"

// ShareWith will return a phrase based on the name passed to the function.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
