package repl

import (
	"bufio"
	"fmt"
	"github.com/cagriyildirimr/ape/lexer"
	"github.com/cagriyildirimr/ape/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
repl:
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Type == token.IDENT && tok.Literal == "exit" {
				break repl
			}
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
