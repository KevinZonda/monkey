package parser

import (
	"github.com/KevinZonda/monkey/ast"
	"github.com/KevinZonda/monkey/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// TODO: parse expression
	for !p.curKindIs(token.SEMICOLON) {
		p.readToken()
	}
	return stmt
}
