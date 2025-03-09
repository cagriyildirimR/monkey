package parser

import (
	"github.com/cagriyildirimr/ape/ast"
	"github.com/cagriyildirimr/ape/lexer"
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
