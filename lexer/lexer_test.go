package lexer

import (
	"go_interpreter/token"
	"testing"
)

func TestPrefixToken(t *testing.T) {
	input := "!0; !foobar;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NOT, "!"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.NOT, "!"},
		{token.IDENT, "foobar"},
	}

	l := New(input)

	for i, tokenTest := range tests {
		tok := l.NextToken()

		if tok.Type != tokenTest.expectedType {
			t.Fatalf("test[%d] - tokentype error. expected=%q got=%q",
				i, tokenTest.expectedType, tok.Type)
		}

		if tok.Literal != tokenTest.expectedLiteral {
			t.Fatalf("test[%d] - tokenliteral error. expected=%q got =%q",
				i, tokenTest.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `def * := \x=0\y. 0
				     | \x\y. + y (* (dec x) y);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.IDENT, "*"},
		{token.ASSIGN, ":="},
		{token.LAMBDA, "\\"},
		{token.IDENT, "x"},
		{token.EQUALITY, "="},
		{token.INT, "0"},
		{token.LAMBDA, "\\"},
		{token.IDENT, "y"},
		{token.DOT, "."},
		{token.INT, "0"},
		{token.OR, "|"},
		{token.LAMBDA, "\\"},
		{token.IDENT, "x"},
		{token.LAMBDA, "\\"},
		{token.IDENT, "y"},
		{token.DOT, "."},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.LPAREN, "("},
		{token.IDENT, "*"},
		{token.LPAREN, "("},
		{token.IDENT, "dec"},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tokenTest := range tests {
		tok := l.NextToken()

		if tok.Type != tokenTest.expectedType {
			t.Fatalf("test[%d] - tokentype error. expected=%q got=%q",
				i, tokenTest.expectedType, tok.Type)
		}

		if tok.Literal != tokenTest.expectedLiteral {
			t.Fatalf("test[%d] - tokenliteral error. expected=%q got =%q",
				i, tokenTest.expectedLiteral, tok.Literal)
		}
	}
}
