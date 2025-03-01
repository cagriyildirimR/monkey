package lexer

import (
	"github.com/cagriyildirimr/ape/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	//input := "=+"
	//
	//tests := []token.Token{
	//	{token.ASSIGN, "="},
	//	{token.PLUS, "+"},
	//}

	tests := []struct {
		input string
		want  []token.Token
	}{
		{
			input: "=+(){},;",
			want: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `let five = 5;
					let ten = 10;
					let add = fn(x, y) {
					x + y;
					};
					let result = add(five, ten);
					`,
			want: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "let 🦊 = 10;",
			want: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "🦊"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "let 变量 = 42;",
			want: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "变量"},
				{token.ASSIGN, "="},
				{token.INT, "42"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "!-/*55;",
			want: []token.Token{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "55"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "5 < 10 > 5;",
			want: []token.Token{
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `if (5 < 10) {
						return true;
					} else {
						return false;
					}`,
			want: []token.Token{
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			input: " 10 == 10;",
			want: []token.Token{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: " 10 != 9;",
			want: []token.Token{
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for i, test := range tests {
		l := New(test.input)
		for j, got := range test.want {
			tok := l.NextToken()
			if tok.Type != got.Type {
				t.Fatalf("tests[%d] - token[%d] - tokentype wrong. expected=%q, got=%q", i, j, got.Type, tok.Type)
			}

			if tok.Literal != got.Literal {
				t.Fatalf("tests[%d] - token[%d] - literal wrong. expected=%q, got=%q", i, j, got.Literal, tok.Literal)
			}
		}
	}

}
