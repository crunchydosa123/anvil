package main

import (
	"fmt"

	"github.com/crunchydosa123/anvil/lexer"
)

func main() {
	input := "let x = 10;"
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
