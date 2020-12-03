// Package proverb generates a proverb.
package proverb

import "fmt"

// Proverb returns a []string proverb using keywords from a provided []string.
func Proverb(rhyme []string) []string {
	var ending = "And all for the want of a %s."
	proverb := make([]string, len(rhyme))

	if len(rhyme) == 0 {
		return proverb
	}

	if len(rhyme) == 1 {
		proverb[0] = fmt.Sprintf(ending, rhyme[0])
	} else {
		for i := 0; i < len(rhyme)-1; i++ {
			proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
		proverb[len(rhyme)-1] = fmt.Sprintf(ending, rhyme[0])
	}

	return proverb
}
