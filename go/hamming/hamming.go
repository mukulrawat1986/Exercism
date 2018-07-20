package hamming

import "errors"

// Distance takes as parameters two strings, and returns the
// hamming distance between the strings and any error.
func Distance(a, b string) (int, error) {
	// if length of strings is not same, return an error
	if len(a) != len(b) {
		return -1, errors.New("strings are not of same length")
	}

	var count int
	// count how many runes are different
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
