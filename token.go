package main

type Token int

const (
	INVALID = iota
	WHITESPACE
	EOF
	IDENTIFIER

	//OPERATORS
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	ASSIGN    // =
	GREATER   // >
	GREATEREQ // >=
	EQUAL     // ==
	NOTEQUAL  // !=
	LESS      // <
	LESSEQ    // <=
	AND       // &&
	OR        // ||

	//SPECIAL CHARACTERS
	COMMA      // ,
	SEMICOL    // ;
	ROUNDLEFT  // (
	ROUNDRIGHT // )
	CURLYLEFT  // {
	CURYRIGHT  //}

	//RESERVED WORDS
	SCANF
	READF
	IF
	ELSE
	WHILE
	INT
	BOOL
	STRUCT
	TRUE
	FALSE
)
