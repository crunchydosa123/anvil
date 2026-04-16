package evaluator

import (
	"fmt"

	"github.com/crunchydosa123/anvil/ast"
)

func Eval(node ast.Node, env *Environment) interface{} {
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
		left := Eval(node.Left, env).(int64)
		right := Eval(node.Right, env).(int64)
		return evalInfix(node.Operator, left, right)

	case *ast.PrintStatement:
		val := Eval(node.Value, env)
		fmt.Println(val)
		return val

	case *ast.FunctionLiteral:
		return &Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        env,
		}

	case *ast.CallExpression:
		function := Eval(node.Function, env).(*Function)
		args := evalExpressions(node.Arguments, env)
		return applyFunction(function, args)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.Value, env)
		return ReturnValue{Value: val}
	}

	return nil
}

func evalProgram(p *ast.Program, env *Environment) interface{} {
	var result interface{}

	for _, stmt := range p.Statements {
		result = Eval(stmt, env)

		if returnVal, ok := result.(ReturnValue); ok {
			return returnVal.Value
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) interface{} {
	var result interface{}

	for _, stmt := range block.Statements {
		result = Eval(stmt, env)

		if _, ok := result.(ReturnValue); ok {
			return result
		}
	}

	return result
}

func applyFunction(fn *Function, args []interface{}) interface{} {
	extendedEnv := NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		extendedEnv.Set(param.Value, args[i])
	}

	result := Eval(fn.Body, extendedEnv)

	return unwrapReturnValue(result)
}

func evalExpressions(exps []ast.Expression, env *Environment) []interface{} {
	var result []interface{}

	for _, e := range exps {
		result = append(result, Eval(e, env))
	}

	return result
}

func unwrapReturnValue(obj interface{}) interface{} {
	if returnVal, ok := obj.(ReturnValue); ok {
		return returnVal.Value
	}
	return obj
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
