package main

import "unicode"

func isWhitespace(sep rune) bool {
	if sep == ' ' || sep == '\n' || sep == '\t' {
		return true
	}
	return false
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
