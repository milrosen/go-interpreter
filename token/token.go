package token

import (
	"slices"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var reservedChars = []byte{'+', ':', '=', '\\', '.', '|', ';', ')', '(', '!'}

func IsReservedChar(ch byte) bool {
	return slices.Contains(reservedChars, ch)
}

var keywords = map[string]TokenType{
	"def": DEF,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = ""

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = ":="
	PLUS     = "+"
	LAMBDA   = "\\"
	EQUALITY = "="
	OR       = "|"
	DOT      = "."
	NOT      = "!"

	// keywords
	DEF = "def"

	// delimiters
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
)
