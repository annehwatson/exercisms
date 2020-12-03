// Package raindrops implements raindrop sound to number conversion.
package raindrops

import (
	"strconv"
)

// Convert returns a corresponding raindrop sound from a number.
func Convert(n int) string {
	result := ""
	if n <= 0 {
		return result
	}

	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}

	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		return strconv.Itoa(n)
	}
	return result
}
