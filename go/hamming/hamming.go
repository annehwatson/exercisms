// Package hamming implements a count of differences in string sequences.
package hamming

import "errors"

// Distance returns the count of differences between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be of equal length")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
