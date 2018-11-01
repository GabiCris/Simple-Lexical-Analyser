package main

import "unicode"

func isWhitespace(sep rune) bool {
	return sep == ' ' || sep == '\t' || sep == '\n' || sep == 13
}

func isLetter(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}
	return false
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}
