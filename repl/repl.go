package repl

import (
	"bufio"
	"fmt"
	"github.com/joeygibson/monkey/lexer"
	"github.com/joeygibson/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.ReadCloser, out io.WriteCloser) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
