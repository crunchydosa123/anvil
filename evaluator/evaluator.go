package evaluator

import "github.com/crunchydosa123/anvil/ast"

func Eval(node ast.Node, env *Environment) int64 {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.Set(node.Name.Value, val)
		return val
	case *ast.IntegerLiteral:
		return node.Value
	case *ast.Identifier:
		val, _ := env.Get(node.Value)
		return val
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfix(node.Operator, left, right)
	}

	return 0
}

func evalProgram(p *ast.Program, env *Environment) int64 {
	var result int64

	for _, stmt := range p.Statements {
		result = Eval(stmt, env)
	}

	return result
}

func evalInfix(op string, left, right int64) int64 {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		return 0
	}
}
