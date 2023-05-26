package parser

import (
	"github.com/KevinZonda/monkey/ast"
	"github.com/KevinZonda/monkey/token"
)

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}

	for p.curToken.Kind != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.readToken()
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Kind {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) expectPeekKind(kind token.TokenKind) bool {
	if p.peekToken.Kind == kind {
		p.readToken()
		return true
	}
	p.peekError(kind)
	return false
}
