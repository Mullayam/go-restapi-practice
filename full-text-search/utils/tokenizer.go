package utils

import (
	"strings"
	"unicode"
)

// tokenize splits the input string into a slice of words, using any non-letter,
// non-number characters as the delimiter. For example, "The quick brown fox"
// would yield the slice ["The", "quick", "brown", "fox"] and "This is a test"
// would yield ["This", "is", "a", "test"].
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
