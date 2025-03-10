package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/claudio4/monkey-interpreter/evaluator"
	"github.com/claudio4/monkey-interpreter/lexer"
	"github.com/claudio4/monkey-interpreter/object"
	"github.com/claudio4/monkey-interpreter/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

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

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			out.Write([]byte{'\n'})
		}

	}
}

func printErrors(out io.Writer, errs []string) {
	for _, err := range errs {
		fmt.Fprintf(out, "\t%s\n", err)
	}
}
