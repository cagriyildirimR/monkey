package parser

import (
	"fmt"
	"github.com/cagriyildirimr/ape/ast"
	"github.com/cagriyildirimr/ape/lexer"
	"github.com/cagriyildirimr/ape/token"
	"strings"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 42;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedExpression string
	}{
		{"x", "5"},
		{"y", "10"},
		{"foobar", "42"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, test.expectedIdentifier, test.expectedExpression) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string, expression string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	if letStmt.Value.TokenLiteral() != expression {
		t.Errorf("letStmt.Value.TokenLiteral() not '%s'. got=%s",
			expression, letStmt.Value.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatementsErrors(t *testing.T) {
	input := `
let x 5;
let  = 10;
let  42;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	tests := []struct {
		want token.Type
		got  token.Type
	}{
		{token.ASSIGN, token.INT},
		{token.IDENT, token.ASSIGN},
		{token.IDENT, token.INT},
	}
	errors := p.Errors()
	if len(errors) != len(tests) {
		t.Fatalf("expected %d errors, got %d", len(tests), len(errors))
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Error %d", i), func(t *testing.T) {
			actual := errors[i]
			if !strings.EqualFold(actual, p.createErrorMessage(test.want, test.got)) {
				t.Errorf("error")
			}
		})
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedExpression string
	}{
		{"5"},
		{"10"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if !testReturnStatement(t, stmt, test.expectedExpression) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement, expression string) bool {
	if s.TokenLiteral() != "return" {
		t.Errorf("s.TokenLiteral not 'return'. got=%q", s.TokenLiteral())
		return false
	}

	returnStmt, ok := s.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("s not *ast.returnStatement. got=%T", s)
		return false
	}

	if returnStmt.Value.TokenLiteral() != expression {
		t.Errorf("returnStmt.Value.TokenLiteral() not '%s'. got=%s",
			expression, returnStmt.Value.TokenLiteral())
		return false
	}

	return true
}
