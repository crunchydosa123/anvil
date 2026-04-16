package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Name  *Identifier
	Value Expression
}

type FunctionLiteral struct {
	Parameters []*Identifier
	Body       *BlockStatement
}

type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

type BlockStatement struct {
	Statements []Statement
}

type ReturnStatement struct {
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return "let"
}

func (ls *LetStatement) String() string {
	var out strings.Builder

	out.WriteString("let ")
	out.WriteString(ls.Name.Value)
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Value
}

type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return ""
}

func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out strings.Builder

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Operator
}

func (ie *InfixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type PrintStatement struct {
	Value Expression
}

func (ps *PrintStatement) statementNode() {}

func (ps *PrintStatement) TokenLiteral() string {
	return "print"
}

func (ps *PrintStatement) String() string {
	var out strings.Builder

	out.WriteString("print(")
	if ps.Value != nil {
		out.WriteString(ps.Value.String())
	}
	out.WriteString(");")

	return out.String()
}

func (fl *FunctionLiteral) expressionNode() {}

func (bs *FunctionLiteral) statementNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return "fn"
}

func (fl *FunctionLiteral) String() string {
	var out strings.Builder

	out.WriteString("fn(")

	for i, p := range fl.Parameters {
		out.WriteString(p.String())
		if i < len(fl.Parameters)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Function.TokenLiteral()
}

func (ce *CallExpression) String() string {
	var out strings.Builder

	out.WriteString(ce.Function.String())
	out.WriteString("(")

	for i, arg := range ce.Arguments {
		out.WriteString(arg.String())
		if i < len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteString(")")

	return out.String()
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return ""
}

func (bs *BlockStatement) String() string {
	var out strings.Builder

	out.WriteString("{ ")

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	out.WriteString(" }")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return "return"
}

func (rs *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString("return ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
