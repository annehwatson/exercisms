// Package accumulate implements a collection modification function.
package accumulate

// Accumulate applies a function to each element in the input collection and
// returns the result in a new collection.
func Accumulate(input []string, converter func(string) string) []string {
	if input == nil {
		return input
	}

	result := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = converter(input[i])
	}
	return result
}
