# Simple Lexical Analyser

Language Specification:

1. Language Definition
1.1.Alphabet:
	a. Upper and lower case letters from the English alphabet;
	b. Decimal digits 0-9.

	Lexic:
	a. Special symbols, representing:
		- operators: + - * / = < <= == != > >= && ||
		- separators: ; , ( ) { }
		- reserved words: scanf readf if else while
						  int bool struct true false

	b. identifiers:
		id ::= letter {(letter | digit)}
		letter ::= "A" | .. | "Z" | "a" | .. | "z"
		digit ::= "0" | .. | "9"
		digitNoZero ::= "1" | .. | "9"

	c. data types:
		- integer:
			integer ::= [+|-] no | "0"
			no ::= digitNoZero{digit}
		- boolean:
			boolean ::= "true"|"false"

1.2 Syntax:
	a. Lexical Rules:
		//ID is any combination of letters and digits
		ID ::= letter {(letter | digit)}
		letter ::= "A" | .. | "Z" | "a" | .. | "z"
		digit ::= "0" | .. | "9"
		RELATION ::= "<" | "<=" | "==" | "!=" | ">" | ">=" | "&&" | "||"
		INTEGER = [+|-] no | "0"
		BOOLEAN = "true"|"false"

	b. Syntactical rules:
		program ::= stmtlist
		stmtlist ::= statement | statement ";" stmtlist
		statement ::= whilestmt | ifstmt | iostmt | assignstmt | struct | "{" {statement} "}"
		whilestmt ::= "while" condition statement
		ifstmt ::= "if" condition statement | "if" condition statement "else" statement
		iostmt ::= "scanf" | "readf" "(" expression ")"
		assignstmt ::= ID "=" expression
		expression ::= expression ("+"|"-") term | term
        term ::= term ("*"|"/") factor | factor
        factor ::= "(" expression ")" | ID | const
		const = INTEGER | BOOLEAN
		condition = expression RELATION expression  
		struct = "struct" "{" {assignstmt} "}"
