package repl

import (
	"bufio"
	"fmt"
	"io"
	"programminglanguage/build"
	"programminglanguage/evaluator"
	"programminglanguage/lexer"
	"programminglanguage/object"
	"programminglanguage/parser"
	"programminglanguage/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer, mode build.Mode) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		if mode == build.TOKENS {
			// Print all tokens and then stop
			tok := l.NextToken()
			for ;tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Println(tok.String())
			}
			continue
		}

		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if mode == build.AST {
			// TODO: ast option not working properly - inverstigate
			fmt.Println(program.String())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
		}

		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
