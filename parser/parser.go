package parser

import (
	"fmt"
	"github.com/cagriyildirimr/ape/ast"
	"github.com/cagriyildirimr/ape/lexer"
	"github.com/cagriyildirimr/ape/token"
)

// Parser is a very simple struct that has
// lexer with current and next token
type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{l: lexer, errors: make([]string, 0)}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()

	for p.currentToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			program.AddStatement(stmt)
		}
		p.NextToken()
	}
	return program
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLet()
	case token.RETURN:
		return p.parseReturn()
	default:
		return nil
	}
}

func (p *Parser) parseReturn() ast.Statement {
	returnToken := p.currentToken

	if !p.expectPeek(token.INT) {
		return nil
	}

	intExp := ast.NewIntExpression(p.currentToken)
	returnStatement := ast.NewReturnStatement(returnToken, intExp)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return returnStatement
}

func (p *Parser) parseLet() ast.Statement {
	letToken := p.currentToken

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	ident := ast.NewIdentifier(p.currentToken)

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.expectPeek(token.INT) { // later add more complex expressions
		return nil
	}
	intExp := ast.NewIntExpression(p.currentToken)
	letStatement := ast.NewLetStatement(letToken, ident, intExp)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return letStatement
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) currentTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	p.errors = append(p.errors, p.createErrorMessage(t, p.peekToken.Type))
}

func (p *Parser) createErrorMessage(got, want token.Type) string {
	return fmt.Sprintf("expected next token to be %s, got %s instead",
		got, want)
}
