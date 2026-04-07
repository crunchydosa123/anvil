package parser

import (
	"strconv"

	"github.com/crunchydosa123/anvil/ast"
	"github.com/crunchydosa123/anvil/lexer"
	"github.com/crunchydosa123/anvil/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
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

	// expect =
	p.nextToken()

	// move to value
	p.nextToken()

	stmt.Value = p.parseIntegerLiteral()
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
