package main

import (
	"fmt"

	"github.com/crunchydosa123/anvil/lexer"
	"github.com/crunchydosa123/anvil/parser"
)

func main() {
	l := lexer.New("let x = 10;")
	p := parser.New(l)

	program := p.ParseProgram()
	fmt.Println(program.String())
}
