/*
Package repl scans the input and outputs all tokens (for now)
*/
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dorin131/dorin-script/lexer"
	"github.com/dorin131/dorin-script/token"
)

const prompt = ">> "

// Start : starts a new REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)
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
