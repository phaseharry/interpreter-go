package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-lang-interpreted/lexer"
	"monkey-lang-interpreted/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	/*
		reads in a new line of text for every loop iteration
		and passes that text to the lexer to generate tokens
	*/
	for {
		fmt.Print(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
