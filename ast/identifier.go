package ast

import "github.com/KevinZonda/monkey/token"

type Identifier struct {
	Token token.Token // the token.IDENTITY token
	Value string
}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i Identifier) expressionNode() {}

var _ Expression = (*Identifier)(nil)
