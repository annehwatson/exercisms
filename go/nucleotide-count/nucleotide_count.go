// Package dna analyzes DNA strands for validity and count of nucleotides.
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(map[byte]int)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	if len(d) == 0 {
		return h, nil
	}

	for i := 0; i < len(d); i++ {
		if _, ok := h[d[i]]; ok {
			h[d[i]]++
		} else {
			return h, errors.New("invalid nucleotide value")
		}
	}

	return h, nil
}
