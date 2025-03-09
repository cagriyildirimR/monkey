package ast

import "github.com/cagriyildirimr/ape/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // convenience method for go compiler to catch if type is statement
}

type Expression interface {
	Node
	expressionNode() // convenience method
}

// Program is a list of Statement
// Program is the root Node of our language
// therefore it needs to implement Node
type Program struct {
	Statements []Statement
}

func NewProgram() *Program {
	return &Program{Statements: make([]Statement, 0)}
}

func (p *Program) AddStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func NewLetStatement(token token.Token, identifier *Identifier, expression Expression) *LetStatement {
	return &LetStatement{
		Token: token,
		Name:  identifier,
		Value: expression,
	}
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// Identifier is an expression not a statement
// even though it doesn't produce a value
// because later we will use some identifier
// in right hand side of the let statement
type Identifier struct {
	Token token.Token
	Value string
}

func NewIdentifier(token token.Token) *Identifier {
	return &Identifier{
		Token: token,
		Value: token.Literal,
	}
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type IntExpression struct {
	Token token.Token
	Value string
}

func (i *IntExpression) expressionNode() {}
func (i *IntExpression) TokenLiteral() string {
	return i.Token.Literal
}

func NewIntExpression(token token.Token) *IntExpression {
	return &IntExpression{
		Token: token,
		Value: token.Literal,
	}
}
