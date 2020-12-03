// Package twofer outputs a simple string.
package twofer

import "fmt"

// ShareWith returns a string message that,
// if a name argument is provided, includes
// the name in the message.
func ShareWith(name string) string {
	message := "One for %s, one for me."

	if name == "" {
		name = "you"
	}

	return fmt.Sprintf(message, name)
}
