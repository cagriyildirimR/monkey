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
		if l.Peek() == '=' {
			var builder strings.Builder
			builder.WriteRune(l.ch)
			l.readRune()
			builder.WriteRune(l.ch)
			tok = newToken(token.EQ, builder.String())
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '!':
		if l.Peek() == '=' {
			var builder strings.Builder
			builder.WriteRune(l.ch)
			l.readRune()
			builder.WriteRune(l.ch)
			tok = newToken(token.NOT_EQ, builder.String())
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
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
			tok = newToken(token.ILLEGAL, string(l.ch))
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

func newToken(ty token.Type, lit string) token.Token {
	return token.Token{
		Type:    ty,
		Literal: lit,
	}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || unicode.IsSymbol(ch)
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

func (l *Lexer) Peek() rune {
	r, _, _ := l.input.ReadRune()
	_ = l.input.UnreadRune()
	return r
}
