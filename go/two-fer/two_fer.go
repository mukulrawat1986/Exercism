package twofer

import "fmt"

// ShareWith takes a string as a parameter and returns a string
// built using the parameter.
func ShareWith(name string) (res string) {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
