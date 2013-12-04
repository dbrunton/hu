package hu

import "strings"

// To "lex" is to separate a sentence into terms.
// The lexer is simple: split on spaces.
func Lex(sentence string) (terms []string, err error) { // []byte ??
	return strings.Split(strings.ToLower(strings.Trim(sentence, ".:,")), " "), nil
}
// I suppose something could have multiple lexings

// To "parse" is to formulate at least one definition of each term to use.
func Parse(terms []string, l Lexicon) (parsed []Definition, err error) {
	return nil, nil
}


