package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tMinamiii/various-parser/monkey/lexer"
	"github.com/tMinamiii/various-parser/monkey/mtoken"
)

const PROMPT = ">> "

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != mtoken.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
