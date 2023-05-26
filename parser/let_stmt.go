package parser

import (
	"github.com/KevinZonda/monkey/ast"
	"github.com/KevinZonda/monkey/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// let <id> = <expr>
	stmt := &ast.LetStatement{
		Token: p.curToken, // let token
	}
	if !p.expectPeekKind(token.IDENTITY) { // <id> token
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	// let <id>
	if !p.expectPeekKind(token.ASSIGN) { // check =
		return nil
	}
	// TODO: parse expression
	for !p.curKindIs(token.SEMICOLON) {
		p.readToken()
	}

	return stmt
}
