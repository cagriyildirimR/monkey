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
	return nil
}
