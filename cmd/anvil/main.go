package main

import (
	"fmt"

	"github.com/crunchydosa123/anvil/evaluator"
	"github.com/crunchydosa123/anvil/lexer"
	"github.com/crunchydosa123/anvil/parser"
)

func main() {
	input := `
	let x = 2 + 3 * 4;
	let y = x + 1;
	`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	env := evaluator.NewEnvironment()
	result := evaluator.Eval(program, env)

	fmt.Println(result)
}
