// Package token is used to define the token type and
// token struct that will be used in the lexer.
package token

// Type represents types in our language
type Type string

type Token struct {
	Type    Type
	Literal string // Literal holds the string representation of the data
}

const (
	ILLEGAL   Type = "ILLEGAL"
	EOF       Type = "EOF"
	IDENT     Type = "IDENT"
	INT       Type = "INT"
	ASSIGN    Type = "="
	PLUS      Type = "+"
	MINUS     Type = "-"
	BANG      Type = "!"
	ASTERISK  Type = "*"
	SLASH     Type = "/"
	EQ        Type = "=="
	NOT_EQ    Type = "!="
	COMMA     Type = ","
	SEMICOLON Type = ";"
	LT        Type = "<"
	GT        Type = ">"
	LPAREN    Type = "("
	RPAREN    Type = ")"
	LBRACE    Type = "{"
	RBRACE    Type = "}"
	FUNCTION  Type = "FUNCTION"
	LET       Type = "LET"
	TRUE      Type = "TRUE"
	FALSE     Type = "FALSE"
	IF        Type = "IF"
	ELSE      Type = "ELSE"
	RETURN    Type = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier checks if given literal is keyword or identifier
func LookupIdentifier(literal string) Type {
	if keywordType, ok := keywords[literal]; ok {
		return keywordType
	}
	return IDENT
}
