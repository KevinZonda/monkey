package parser

import (
	"github.com/KevinZonda/monkey/lexer"
	"github.com/KevinZonda/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func (p *Parser) curKindIs(kind token.TokenKind) bool {
	return p.curToken.Kind == kind
}

func (p *Parser) peekKindIs(kind token.TokenKind) bool {
	return p.peekToken.Kind == kind
}


func (p *Parser) readToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.readToken()
	p.readToken()
	return p
}
