// Package token is used to define the token type and token struct that will be used in the lexer.
package token

// Type is the type of data in source code
type Type string

type Token struct {
	Type    Type
	Literal string // Literal holds the string representation of the data
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdentifier checks if given literal is keyword or identifier
func LookupIdentifier(literal string) Type {
	if keywordType, ok := keywords[literal]; ok {
		return keywordType
	}
	return IDENT
}
