package evaluator

import "github.com/crunchydosa123/anvil/ast"

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

type ReturnValue struct {
	Value interface{}
}
