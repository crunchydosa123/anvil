package main

import (
	"fmt"
	"os"

	"github.com/crunchydosa123/anvil/evaluator"
	"github.com/crunchydosa123/anvil/lexer"
	"github.com/crunchydosa123/anvil/parser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: anvil run <file>")
		return
	}

	command := os.Args[1]
	filename := os.Args[2]

	switch command {
	case "run":
		runFile(filename)
	default:
		fmt.Println("unknown command:", command)
	}
}

func runFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	input := string(content)

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Println(" -", msg)
		}
		return
	}

	env := evaluator.NewEnvironment()
	evaluator.Eval(program, env)
}
