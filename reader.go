package main

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	reader *bufio.Reader
}

// Wrapper for the reader buffer class
func getScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader)}
}

func (scanner *Scanner) read() rune {
	char, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return char
}

func (scanner *Scanner) unread() {
	_ = scanner.reader.UnreadRune()
}

func checkKeyword(str string) (Token, string) {
	switch strings.ToUpper(str) {
	case "SCANF":
		return SCANF, str
	case "READF":
		return READF, str
	case "IF":
		return IF, str
	case "ELSE":
		return ELSE, str
	case "WHILE":
		return WHILE, str
	case "INT":
		return INT, str
	case "BOOL":
		return BOOL, str
	case "STRUCT":
		return STRUCT, str
	case "TRUE":
		return TRUE, str
	case "FALSE":
		return FALSE, str
	}
	return 0, ""
}

func tokenizeOperator(char rune, scanner *Scanner) (Token, string) {
	switch char {
	case eof:
		return EOF, ""
	case '/':
		return DIV, string(char)
	case '*':
		return MUL, string(char)
	case ',':
		return COMMA, string(char)
	case ';':
		return SEMICOL, string(char)
	case '(':
		return ROUNDLEFT, string(char)
	case ')':
		return ROUNDRIGHT, string(char)
	case '{':
		return CURLYLEFT, string(char)
	case '}':
		return CURYRIGHT, string(char)
	case '-':
		return ADD, string(char)
	case '+':
		return SUB, string(char)
	case '=':
		if charNext := scanner.read(); charNext == '=' {
			return EQUAL, "=="
		} else {
			scanner.unread()
			return ASSIGN, string(char)
		}
	case '>':
		if charNext := scanner.read(); charNext == '=' {
			return GREATEREQ, ">="
		} else {
			scanner.unread()
			return GREATER, string(char)
		}
	case '<':
		if charNext := scanner.read(); charNext == '=' {
			return LESSEQ, "<="
		} else {
			scanner.unread()
			return LESS, string(char)
		}
	case '!':
		if charNext := scanner.read(); charNext == '=' {
			return NOTEQUAL, "!="
		} else {
			return INVALID, string(char)
		}
	case '&':
		if charNext := scanner.read(); charNext == '&' {
			return AND, "&&"
		} else {
			return INVALID, string(char)
		}
	case '|':
		if charNext := scanner.read(); charNext == '|' {
			return OR, "||"
		} else {
			return INVALID, string(char)
		}
	}
	return INVALID, string(char)
}
