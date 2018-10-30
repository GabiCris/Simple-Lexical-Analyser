package main

import "io"

type Parser struct {
	scanner *Scanner
	buffer  struct {
		token Token
		str   string
		size  int
	}
}

func getParser(reader io.Reader) *Parser {
	return &Parser{scanner: getScanner(reader)}
}

// scan and unscan for parser wrapper
func (parser *Parser) scanParser() (token Token, str string) {
	if parser.buffer.size != 0 {
		parser.buffer.size = 0
		return parser.buffer.token, parser.buffer.str
	}

	token, str = parser.scanner.scan()

	parser.buffer.token, parser.buffer.str = token, str
	return
}

func (parser *Parser) unscanParser() {
	parser.buffer.size = 1
}

func (parser *Parser) scanWithoutSpaces() (token Token, str string) {
	token, str = parser.scanParser()
	if token == WHITESPACE {
		token, str = parser.scanParser()
	}
	return
}
