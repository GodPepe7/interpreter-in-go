package repl

import (
	"bufio"
	"fmt"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
	"io"
)

/* REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
/  interpreter and starts again. Read, Eval, Print, Loop.
*/

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
