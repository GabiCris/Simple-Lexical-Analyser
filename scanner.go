package main

import (
	"bytes"
)

var eof = rune(0)

// Scans for multipe CHARACTERS
// continuous Whitespace scan
func (scanner *Scanner) scanWhitespace() (token Token, lit string) {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.read())
	for {
		if char := scanner.read(); char == eof {
			break
		} else if !isWhitespace(char) {
			scanner.unread()
			break
		} else {
			buffer.WriteRune(char)
		}
	}
	return WHITESPACE, buffer.String()
}

// identifier scan
func (scanner *Scanner) scanIdentifier() (token Token, lit string) {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.read())

	for {
		if char := scanner.read(); char == eof {
			break
		} else if !isLetter(char) && !isDigit(char) {
			scanner.unread()
			break
		} else {
			_, _ = buffer.WriteRune(char)
		}
	}
	keyWordToken, str := checkKeyword(buffer.String())

	if keyWordToken == 0 && str == "" {
		return IDENTIFIER, buffer.String()
	}

	return keyWordToken, str
}

func (scanner *Scanner) scanNumberConstant() (token Token, lit string) {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.read())

	for {
		if char := scanner.read(); char == eof {
			break
		} else if !isDigit(char) {
			scanner.unread()
			break
		} else {
			buffer.WriteRune(char)
		}
	}
	keyword, str := verifyNumberConstant(buffer.String())
	return keyword, str
}

// scan function which will return the next token and the respective string
func (scanner *Scanner) scan() (token Token, lit string) {
	char := scanner.read()

	// Tokenizing whitespaces and identifiers
	if isWhitespace(char) {
		scanner.unread()
		return scanner.scanWhitespace()
	} else if isLetter(char) {
		scanner.unread()
		return scanner.scanIdentifier()
	} else if isDigit(char) {
		scanner.unread()
		return scanner.scanNumberConstant()
	}

	// otherwise tokenizing operators; if not an operator returns INVALID
	return tokenizeOperator(char, scanner)
}
