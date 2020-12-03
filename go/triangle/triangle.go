// Package triangle implements triangle classification.
package triangle

import "math"

// Kind represents the type of triangle.
type Kind string

const (
	// NaT represents Not a Triangle.
	NaT = "NaT"
	// Equ represents an Equilateral triangle.
	Equ = "Equ"
	// Iso represents an Isoceles triangle.
	Iso = "Iso"
	// Sca represents a Scalene triangle.
	Sca = "Sca"
)

// KindFromSides returns a triangle's type.
func KindFromSides(a, b, c float64) Kind {
	if !IsAValidTriangle(a, b, c) {
		return NaT
	}

	set := make(map[float64]bool)
	set[a] = true
	set[b] = true
	set[c] = true

	count := len(set)

	if count == 1 {
		return Equ
	}
	if count == 2 {
		return Iso
	}
	if count == 3 {
		return Sca
	}

	return NaT
}

// IsAValidTriangle checks for triangle validity and returns true if a valid
// triangle or false if not a valid triangle.
func IsAValidTriangle(a, b, c float64) bool {
	if a < 0 || b < 0 || c < 0 {
		return false
	}

	if a+b < c || a+c < b || c+b < a {
		return false
	}

	if a == 0 && b == 0 && c == 0 {
		return false
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if a == math.Inf(-1) || b == math.Inf(-1) || c == math.Inf(-1) {
		return false
	}

	if a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1) {
		return false
	}
	return true
}
