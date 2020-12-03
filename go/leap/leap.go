// Package leap implements leap year evaluation.
package leap

// IsLeapYear returns true if the argument is a leap year, otherwise false.
func IsLeapYear(int) bool {
	return (int % 400 == 0 || int % 4 == 0 && int % 100 != 0)
}
