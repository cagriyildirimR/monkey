package parser

import (
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
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{l: lexer}

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

	for !(p.currentToken.Type == token.SEMICOLON && p.peekToken.Type == token.EOF) {
		if p.currentToken.Type == token.LET {
			program.AddStatement(p.parseLet())
		} else {
			p.NextToken()
		}
	}
	return program
}

func (p *Parser) parseLet() ast.Statement {
	letToken := p.currentToken
	if p.peekToken.Type != token.IDENT {
		return nil
	}
	ident := ast.NewIdentifier(p.peekToken)
	p.NextToken() // current token is ident next token is equal
	if p.peekToken.Type != token.ASSIGN {
		return nil
	}
	p.NextToken() // expr
	intExp := ast.NewIntExpression(p.peekToken)
	letStatement := ast.NewLetStatement(letToken, ident, intExp)
	return letStatement
}
