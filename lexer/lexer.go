package lexer

import (
	"github.com/cagriyildirimr/ape/token"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	input *strings.Reader
	ch    rune
}

func New(input string) *Lexer {
	lexer := &Lexer{input: strings.NewReader(input)}
	lexer.readRune()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readRune()
	return tok
}

func (l *Lexer) readRune() {
	r, _, err := l.input.ReadRune()
	if err == io.EOF {
		l.ch = 0
	} else {
		l.ch = r
	}
}

func newToken(ty token.Type, lit rune) token.Token {
	return token.Token{
		Type:    ty,
		Literal: string(lit),
	}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch == '-' || unicode.IsSymbol(ch)
}

func (l *Lexer) readIdentifier() string {
	var builder strings.Builder
	for isLetter(l.ch) {
		builder.WriteRune(l.ch)
		l.readRune()
	}
	return builder.String()
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readRune()
	}
}

func (l *Lexer) readNumber() string {
	var builder strings.Builder
	for isDigit(l.ch) {
		builder.WriteRune(l.ch)
		l.readRune()
	}
	return builder.String()
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
