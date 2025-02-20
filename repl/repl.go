package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/claudio4/going-monkey/lexer"
	"github.com/claudio4/going-monkey/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		out.Write([]byte{'\n'})
	}
}

func printErrors(out io.Writer, errs []string) {
	for _, err := range errs {
		fmt.Fprintf(out, "\t%s\n", err)
	}
}
