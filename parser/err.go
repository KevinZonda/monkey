package parser

import (
	"fmt"
	"github.com/KevinZonda/monkey/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenKind) {
	msg := "expected next token %s, got %s"
	p.errors = append(p.errors, fmt.Sprintf(msg, t, p.peekToken.Kind))
}
