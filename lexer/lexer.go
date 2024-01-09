package lexer

import (
	"go_interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // position of the beginning of a meaningfull phrase, [pos]let
	readPosition int  // posiiton of the reading head,                      le[pos]t
	ch           byte // current char being read
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// moves one forward, does not return a value
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// returns next character, does not read
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '|':
		tok = newToken(token.OR, l.ch)
		l.readChar()
	case '!':
		tok = newToken(token.NOT, l.ch)
		l.readChar()
	case '+':
		tok = newToken(token.PLUS, l.ch)
		l.readChar()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case '.':
		tok = newToken(token.DOT, l.ch)
		l.readChar()
	case '\\':
		tok = newToken(token.LAMBDA, l.ch)
		l.readChar()
	case '=':
		tok = newToken(token.EQUALITY, l.ch)
		l.readChar()
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.readChar()
	case ':':
		l.readChar()
		if l.ch == '=' {
			tok.Literal = ":="
			tok.Type = token.ASSIGN
			l.readChar()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isIdentifier(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.INT
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	return tok
}

func (l *Lexer) readDigit() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isIdentifier(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isIdentifier(ch byte) bool {
	return isPrintable(ch) && !isDigit(ch) && !token.IsReservedChar(ch)
}

func isDigit(ch byte) bool {
	// 48-57 is 0-9 in ascii
	return ch >= 48 && ch <= 57
}

func isPrintable(ch byte) bool {
	// 33-126 is all printable ascii characters as bytes. Later we check if they are just the byte.
	// this also ignores all whitespace!
	return ch >= 33 && ch <= 126
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' || l.ch == '\v' {
		l.readChar()
	}
}
