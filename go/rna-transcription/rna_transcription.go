// Package strand implements DNA to RNA transcription.
package strand

// ToRNA takes in a DNA string and returns its RNA complement.
func ToRNA(dna string) string {
	rna := ""

	dnaToRna := make(map[byte]string)
	dnaToRna['G'] = "C"
	dnaToRna['C'] = "G"
	dnaToRna['T'] = "A"
	dnaToRna['A'] = "U"

	for i := 0; i < len(dna); i++ {
		rna += dnaToRna[dna[i]]
	}

	return rna
}
