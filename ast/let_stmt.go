package ast

import "github.com/KevinZonda/monkey/token"

// Let -> let <identifier> = <expression>

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (l LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l LetStatement) statementNode() {}

var _ Statement = (*LetStatement)(nil)
