package parser

import (
	"fmt"
	"strconv"

	"github.com/crunchydosa123/anvil/ast"
	"github.com/crunchydosa123/anvil/lexer"
	"github.com/crunchydosa123/anvil/token"
)

const (
	_ int = iota
	LOWEST
	SUM
	PRODUCT
	CALL
)

var precedences = map[token.Type]int{
	token.PLUS:     SUM,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.PRINT:
		return p.parsePrintStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{}

	// expect identifier
	p.nextToken()

	stmt.Name = &ast.Identifier{
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// move to value
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, _ := strconv.ParseInt(p.curToken.Literal, 0, 64)

	return &ast.IntegerLiteral{
		Value: val,
	}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	left := p.parsePrimary()

	for p.peekToken.Type != token.SEMICOLON && precedence < p.peekPrecedence() {
		p.nextToken()

		switch p.curToken.Type {
		case token.PLUS, token.ASTERISK:
			left = p.parseInfixExpression(left)
		case token.LPAREN:
			left = p.parseCallExpression(left)
		}

	}
	return left
}

func (p *Parser) parsePrimary() ast.Expression {
	switch p.curToken.Type {
	case token.INT:
		return p.parseIntegerLiteral()
	case token.IDENT:
		return &ast.Identifier{Value: p.curToken.Literal}
	case token.LPAREN:
		p.nextToken()
		exp := p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
		return exp
	case token.FUNCTIONDEF:
		return p.parseFunction()

	default:
		return nil
	}
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Left:     left,
		Operator: p.curToken.Literal,
	}

	precedence := p.currPrecedence()
	p.nextToken()

	exp.Right = p.parseExpression(precedence)

	return exp
}

func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{}

	p.nextToken()

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
	}

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	msg := fmt.Sprintf(
		"expected token to be %s, got %s instead",
		t,
		p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)

	return false
}

func (p *Parser) parseFunction() ast.Expression {
	lit := &ast.FunctionLiteral{}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{
		Function: function,
	}

	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var params []*ast.Identifier

	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
		return params
	}

	p.nextToken()
	params = append(params, &ast.Identifier{
		Value: p.curToken.Literal,
	})

	for p.peekToken.Type == token.COMMA {
		p.nextToken()
		p.nextToken()

		params = append(params, &ast.Identifier{
			Value: p.curToken.Literal,
		})

		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	}

	return params
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var args []ast.Expression

	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekToken.Type == token.COMMA {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}
