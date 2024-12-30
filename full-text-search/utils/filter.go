package utils

import (
	snowballeng "github.com/kljensen/snowball/english"
	"strings"
)

func lowercaseFilter(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}
	return tokens
}

func stopwordFilter(tokens []string) []string {

	var stopwords = map[string]struct{}{
		"a":    {},
		"an":   {},
		"and":  {},
		"are":  {},
		"as":   {},
		"at":   {},
		"be":   {},
		"by":   {},
		"for":  {},
		"from": {},
		"has":  {},
		"he":   {},
		"in":   {},
		"is":   {},
		"it":   {},
		"its":  {},
		"of":   {},
		"on":   {},
		"that": {},
		"the":  {},
		"to":   {},
		"was":  {},
		"were": {},
		"will": {},
		"with": {},
	}
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}

	}
	return tokens
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
