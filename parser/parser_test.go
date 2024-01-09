package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"testing"
)

func TestPrefixExpression(t *testing.T) {
	input := "!0;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("expected one statement only, recieved %d instead", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatement, got %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.PrefixExpression)

	if !ok {
		t.Fatalf("program.Statements[0] is not a prefix, got %T", stmt.Expression)
	}

	if exp.Operator != "!" {
		t.Fatalf("Identifier operator not '!', got %s instead", exp.Operator)
	}

	integ, ok := exp.Right.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("did not recieve integer literal as right end, got %T instead", exp.Right)
	}

	if integ.Value != 0 {
		t.Fatalf("did not get 0 as value for !0, instead recieved %d", integ.Value)
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("expected one statement only, recieved %d instead", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("program.Statements[0] is not an Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("Identifier value not foobar, got %s instead", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("Expression Literal value not foobar, got %s instead", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("expected one statement only, recieved %d instead", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("program.Statements[0] is not an Identifier, got %T", stmt.Expression)
	}

	if ident.Value != 5 {
		t.Fatalf("Identifier value not foobar, got %d instead", ident.Value)
	}

	if ident.TokenLiteral() != "5" {
		t.Fatalf("Expression Literal value not foobar, got %s instead", ident.TokenLiteral())
	}
}

func TestDefErrorStatement(t *testing.T) {
	input := `def :=;`

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()

	if len(p.Errors()) != 2 {
		t.Fatalf("Missing Identifier Not Detected")
	}

	if p.Errors()[0] != "expected next token to be IDENT, got := instead" {
		t.Fatalf("Expected Ident error, got {%s} instead", p.Errors()[0])
	}
}

func TestDefStatements(t *testing.T) {
	input := `def pi := 3; def four := 5;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("Expected 2 statements, recieved %d instead", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"pi"},
		{"four"},
	}

	for i, name := range tests {
		statement := program.Statements[i]
		if !testDefStatement(t, statement, name.expectedIdentifier) {
			return
		}
	}
}

func testDefStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "def" {
		t.Errorf("statement.TokenLiteral %q instead of 'def'", statement.TokenLiteral())
		return false
	}

	defStatement, ok := statement.(*ast.DefStatement)

	if !ok {
		t.Errorf("statement not DefStatement, got %T", statement)
		return false
	}

	if defStatement.Name.Value != name {
		t.Errorf("defStatement.Name.Value not %s, recieved %s instead", name, defStatement.Name.Value)
		return false
	}

	if defStatement.Name.TokenLiteral() != name {
		t.Errorf("defStatement.Name.TokenLiteral not %s, recieved %s instead", name, defStatement.Name.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
